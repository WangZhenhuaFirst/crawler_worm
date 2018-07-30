package main

import (
	"testing"
	"time"

	"github.com/crawler/engine"
	"github.com/crawler/model"
	"github.com/crawler_distributed/config"
	"github.com/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
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

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}

}
