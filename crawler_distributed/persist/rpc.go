package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"learngo_base/crawler/engine"
	"learngo_base/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
