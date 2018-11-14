package main

import (
	"learngo_base/crawler/engine"
	"learngo_base/crawler/model"
	"learngo_base/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemsaver(t *testing.T) {
	const host = ":1234"
	//  start ItemSaverServer
	go serverRpc(host, "test1")
	time.Sleep(time.Second)

	//	start ItemServerClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//  call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	result := ""
	err = client.Call("&ItemSaverService.Save",
		//err = client.Call(config.ItemSaverRpc,
		item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s",
			result, err)
	}
}
