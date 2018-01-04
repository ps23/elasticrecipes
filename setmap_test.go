package elasticrecipes

import (
	"log"
	"testing"

	elastic "github.com/olivere/elastic"
)

func TestSetMap(t *testing.T) {
	type args struct {
		client      *elastic.Client
		index       string
		mapjsonfile string
	}

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// Handle error
		log.Println(err.Error())
	}
	tests := []struct {
		name string
		args args
	}{
		{"test0", args{client, "elastic-test", "resources/mapping.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMap(tt.args.client, tt.args.index, tt.args.mapjsonfile)
		})
	}
}
