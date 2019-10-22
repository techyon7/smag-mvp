package scraper_client

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/codeuniversity/smag-mvp/awsService/proto"
	"github.com/codeuniversity/smag-mvp/http-header-generator"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var userAccountInfoUrl = "https://instagram.com/%s/?__a=1"
var userAccountMediaUrl = "https://www.instagram.com/graphql/query/?query_hash=58b6785bea111c67129decbe6a448951&variables=%s"
var userPostsCommentUrl = "https://www.instagram.com/graphql/query/?query_hash=865589822932d1b43dfe312121dd353a&variables=%s"

type HttpClient struct {
	*http_header_generator.HttpHeaderGenerator
	localAddressReachLimit bool
	localIp                string
	currentAddress         string
	client                 *http.Client
	instanceId             string
	grpcClient             *grpc.ClientConn
}

func NewHttpClient(awsServiceAddress string) *HttpClient {
	client := &HttpClient{}
	client.HttpHeaderGenerator = http_header_generator.New()
	var err error

	addresses := getLocalIpAddresses(1)
	//addresses := []string{"192.168.178.41"}

	client.localIp = addresses[0]

	client.client, err = client.getBoundAddressClient(addresses[0])

	if err != nil {
		panic(err)
	}

	client.instanceId = getAmazonInstanceId()

	client.grpcClient, err = grpc.Dial(awsServiceAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return client
}

func getAmazonInstanceId() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	return string(body)
}

func isNetworkInterfaces(name string) bool {
	matched, err := regexp.MatchString("eth[0-9]*$", name)

	if err != nil {
		panic(err)
	}

	return matched
}

func isIpv4Address(ip string) bool {
	matched, err := regexp.MatchString("^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$", ip)

	if err != nil {
		panic(err)
	}

	return matched
}

func getLocalIpAddresses(count int) []string {
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Println("Get Network Interfaces Error: ")
		panic(err)
	}

	var localAddresses []string
	for _, networkInterface := range interfaces {
		if isNetworkInterfaces(networkInterface.Name) {
			addrs, err := networkInterface.Addrs()
			if err != nil {
				fmt.Println("Error Addrs: ", err)
				panic(err)
			}
			for _, address := range addrs {
				ip := strings.Split(address.String(), "/")
				if isIpv4Address(ip[0]) {
					localAddresses = append(localAddresses, ip[0])
				}
			}
		}
	}

	if len(localAddresses) < count {
		panic(fmt.Sprintf("Not Enough Local Ip Addresses, Requirement: %d \n", count))
	}

	fmt.Println("All LocalAddresses: ", localAddresses)
	return localAddresses[:count]
}

func (h *HttpClient) getBoundAddressClient(localIp string) (*http.Client, error) {
	localAddr, err := net.ResolveIPAddr("ip", localIp)

	if err != nil {
		return nil, err
	}

	localTCPAddr := net.TCPAddr{
		IP: localAddr.IP,
	}

	d := net.Dialer{
		LocalAddr: &localTCPAddr,
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	tr := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         d.DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return &http.Client{Transport: tr}, nil
}

func (h *HttpClient) WithRetries(times int, f func() error) error {
	var err error
	for i := 0; i < times; i++ {
		err = f()
		if err == nil {
			return nil
		}

		fmt.Println(err)
		foundAddress, err := h.checkIfIPReachedTheLimit(err)
		fmt.Println("FoundAddress: ", foundAddress)
		if err != nil {
			fmt.Println(err)
		}
		if foundAddress {
			times++
		}
		time.Sleep(100 * time.Millisecond)
	}
	return err
}

func (h *HttpClient) checkIfIPReachedTheLimit(err error) (bool, error) {
	fmt.Println("checkIfIPReachedTheLimit")
	switch t := err.(type) {
	case *json.SyntaxError:
		fmt.Println("SyntaxError")

		if h.localAddressReachLimit == true {
			_, err := h.sendRenewElasticIpRequestToAmazonService()
			if err != nil {
				return false, err
			}

			h.localAddressReachLimit = false
			return true, nil

		}
	case *json.UnmarshalTypeError:
		fmt.Println("UnmarshalTypeError")
	case *json.InvalidUnmarshalError:
		fmt.Println("InvalidUnmarchedError")
	case *json.UnsupportedTypeError:
		fmt.Println("UnsupportedTypeError")
	case *HttpStatusError:
		if h.localAddressReachLimit == true {
			_, err := h.sendRenewElasticIpRequestToAmazonService()
			if err != nil {
				return false, err
			}

			h.localAddressReachLimit = false
			return true, nil

		}
	default:
		fmt.Println("Found Wrong Json Type Error ", t)
		return false, err
	}
	fmt.Println("checkIfIPReachedTheLimit is not working!!!")
	return false, err
}

//func (h *HttpClient) checkAvailableAddresses() ([]string, bool) {
//	h.localAddressesReachLimit[h.currentAddress] = false
//	var addresses []string
//	var err error
//	for ip := range h.localAddressesReachLimit {
//		addresses = append(addresses, ip)
//		if h.localAddressesReachLimit[ip] == true {
//			h.currentAddress = ip
//			h.client, err = h.getBoundAddressClient(ip)
//			if err != nil {
//				panic(err)
//			}
//			fmt.Println("Update Client")
//			return addresses, true
//		}
//	}
//	return addresses, false
//}
func (h *HttpClient) sendRenewElasticIpRequestToAmazonService() (bool, error) {

	renewIp := pb.RenewingElasticIp{
		InstanceId: h.instanceId,
		Node:       "",
		Pod:        "",
		PodIp:      h.localIp,
	}

	awsClient := pb.NewRouteGuideClient(h.grpcClient)
	result, err := awsClient.RenewElasticIp(context.Background(), &renewIp)
	if err != nil {
		fmt.Println("sendRenewElasticIpRequestToAmazonService Err: ", err)
		return false, err
	}
	return result.IsRenewed, nil
}

//func (h *HttpClient) waitForRenewElasticIpRequest() (*models.RenewingAddresses, error) {
//	fmt.Println("waitForRenewElasticIpRequest")
//	message, err := h.renewedAddressQReader.FetchMessage(context.Background())
//	fmt.Println("waitForRenewElasticIpRequest Finished: ")
//	if err != nil {
//		fmt.Println("waitForRenewElasticIpRequest error")
//		return nil, err
//	}
//	fmt.Println("Wait Message Time: ", message.Time)
//
//	var renewedAddresses models.RenewingAddresses
//	err = json.Unmarshal(message.Value, &renewedAddresses)
//	if err != nil {
//		return nil, err
//	}
//
//	h.renewedAddressQReader.CommitMessages(context.Background(), message)
//	return &renewedAddresses, err
//}
//

func (h *HttpClient) Close() {
	h.grpcClient.Close()
}

func (h *HttpClient) Do(request *http.Request) (*http.Response, error) {
	h.AddHeaders(&request.Header)
	return h.client.Do(request)
}
