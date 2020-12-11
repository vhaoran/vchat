package yes

import (
	"context"

	elastic "github.com/olivere/elastic/v7"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	X *elastic.Client
)

func InitES(cfg yconfig.ESConfig) error {
	conn, err := NewESCnt(cfg)
	if err != nil {
		return err
	}
	X = conn
	return nil
}

//初始化
func NewESCnt(cfg yconfig.ESConfig) (*elastic.Client, error) {
	var client *elastic.Client
	var err error
	client, err = elastic.NewClient(elastic.SetURL(cfg.Url...),
		elastic.SetSniff(cfg.Sniff))
	if err != nil {
		return nil, err
	}

	for _, url := range cfg.Url {
		_, _, err = client.Ping(url).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}

	for _, url := range cfg.Url {
		_, err = client.ElasticsearchVersion(url)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}
