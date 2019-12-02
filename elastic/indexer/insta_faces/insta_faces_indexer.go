package main

import (
	"encoding/json"
	"strconv"

	"github.com/codeuniversity/smag-mvp/elastic"
	"github.com/codeuniversity/smag-mvp/elastic/indexer"
	esModels "github.com/codeuniversity/smag-mvp/elastic/models"
	"github.com/codeuniversity/smag-mvp/kafka/changestream"
	"github.com/codeuniversity/smag-mvp/models"
	"github.com/codeuniversity/smag-mvp/service"
	"github.com/codeuniversity/smag-mvp/utils"
)

func main() {
	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "my-kafka:9092")
	groupID := utils.MustGetStringFromEnv("KAFKA_GROUPID")
	bulkSize := utils.GetNumberFromEnvWithDefault("BULK_SIZE", 10)
	changesTopic := utils.GetStringFromEnvWithDefault("KAFKA_CHANGE_TOPIC", "postgres.public.face_data")

	esHosts := utils.GetMultipleStringsFromEnvWithDefault("ES_HOSTS", []string{"http://localhost:9200"})

	i := indexer.New(esHosts, elastic.FacesIndex, elastic.FacesIndexMapping, kafkaAddress, changesTopic, groupID, indexFace, bulkSize)

	service.CloseOnSignal(i)
	waitUntilDone := i.Start()

	waitUntilDone()
}

func indexFace(m *changestream.ChangeMessage) (*indexer.ElasticIndexer, error) {

	switch m.Payload.Op {
	case "r", "u", "c":
		break
	default:
		return &indexer.ElasticIndexer{}, nil
	}

	face := &models.FaceData{}
	err := json.Unmarshal(m.Payload.After, face)
	if err != nil {
		return &indexer.ElasticIndexer{}, err
	}

	return createBulkIndexOperation(face)
}

func createBulkIndexOperation(face *models.FaceData) (*indexer.ElasticIndexer, error) {
	bulkOperation := `{ "index": {}  }`

	bulkOperationJson, err := json.Marshal(bulkOperation)

	if err != nil {
		return &indexer.ElasticIndexer{}, err
	}

	bulkOperationJson = append(bulkOperationJson, "\n"...)

	doc, err := esModels.FaceDocFromFaceData(face)
	if err != nil {
		return &indexer.ElasticIndexer{}, err
	}

	docJson, err := json.Marshal(doc)

	if err != nil {
		return &indexer.ElasticIndexer{}, err
	}

	docJson = append(docJson, "\n"...)

	bulkUpsertBody := string(bulkOperationJson) + string(docJson)

	return &indexer.ElasticIndexer{DocumentId: strconv.Itoa(int(face.ID)), BulkOperation: bulkUpsertBody}, err

}
