package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/crawler_distributed/config"
	"github.com/crawler_distributed/rpcsupport"
	"github.com/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1826404337",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "vivian",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
