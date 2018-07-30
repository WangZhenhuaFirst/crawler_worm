package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/crawler/engine"
	"github.com/crawler/model"
	elastic "gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1826404337",
		Type: "zhenai",
		Id:   "1826404337",
		Payload: model.Profile{
			Age:        26,
			Height:     168,
			Weight:     53,
			Income:     "3000元以下",
			Gender:     "女",
			Name:       "vivian",
			Xinzuo:     "魔羯座",
			Occupation: "咨询/顾问",
			Marriage:   "未婚",
			House:      "租房",
			Hokou:      "山东济南",
			Education:  "硕士",
			Car:        "未购车",
		},
	}

	//TODO: Try to start up elastic search here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
