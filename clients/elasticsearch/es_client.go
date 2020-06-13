package elasticsearch

import (
	"time"

	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

// variable ecualt to a pointer of elasticSearch client
var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

// Init func
func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.1.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}
