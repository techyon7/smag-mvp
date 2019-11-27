package main

import (
	"encoding/json"
	"fmt"
	"github.com/codeuniversity/smag-mvp/elastic"
	es_indexer "github.com/codeuniversity/smag-mvp/elastic/indexer"
	"github.com/codeuniversity/smag-mvp/kafka/changestream"
	"github.com/codeuniversity/smag-mvp/service"
	"github.com/codeuniversity/smag-mvp/utils"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"strconv"
)

func main() {
	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "my-kafka:9092")
	groupID := utils.MustGetStringFromEnv("KAFKA_GROUPID")
	changesTopic := utils.GetStringFromEnvWithDefault("KAFKA_CHANGE_TOPIC", "postgres.public.posts")

	esHosts := utils.GetMultipliesStringsFromEnvDefault("ELASTIC_SEARCH_ADDRESS", []string{"localhost:9201"})

	elasticInserter := es_indexer.New(esHosts, elastic.CommentsIndex, elastic.CommentsIndexMapping, kafkaAddress, changesTopic, groupID, indexComment)

	service.CloseOnSignal(elasticInserter)
	waitUntilClosed := elasticInserter.Start()

	waitUntilClosed()
}

func indexComment(client *elasticsearch.Client, m *changestream.ChangeMessage) error {
	comment := &comment{}
	err := json.Unmarshal(m.Payload.After, comment)

	if err != nil {
		return err
	}

	switch m.Payload.Op {
	case "r", "c":
		return upsertComment(comment, client)
	}

	return nil
}

func upsertComment(comment *comment, client *elasticsearch.Client) error {

	upsertBody := createUpsertBody(comment)
	response, err := client.Update(elastic.CommentsIndex, strconv.Itoa(comment.ID), esutil.NewJSONReader(upsertBody))

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("upsertDocument Upsert Document Failed StatusCode=%s Body=%s", response.Status(), response.String())
	}
	return nil
}

func createUpsertBody(comment *comment) map[string]interface{} {
	var commentUpsert = map[string]interface{}{
		"script": map[string]interface{}{
			"source": "ctx._source.comment = params.comment",
			"lang":   "painless",
			"params": map[string]interface{}{
				"comment": comment.Comment,
			},
		},
		"upsert": map[string]interface{}{
			"post_id": comment.PostID,
			"comment": comment.Comment,
		},
	}

	return commentUpsert
}

type comment struct {
	ID      int    `json:"id"`
	PostID  string `json:"post_id"`
	Comment string `json:"comment_text"`
}
