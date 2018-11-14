package main

import (
	"gopkg.in/olivere/elastic.v5"
	"learngo_base/crawler_distributed/persist"
	"learngo_base/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	//强制退出..
	log.Fatal(serverRpc(":1234", "dating_profile"))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
