package elasticrecipes

import (
	"context"
	"io/ioutil"
	"log"

	elastic "github.com/olivere/elastic"
)

// Create Index and config trough mapjsonfile
func SetMap(client *elastic.Client, index string, mapjsonfile string) (created bool, e error) {
	var err error
	var indexcreated bool

	log.Println("setting ", mapjsonfile, " for index ", index)

	// Read in nginx_json_template
	buf, err := ioutil.ReadFile(mapjsonfile)
	if err != nil {
		log.Fatal(err)
		return indexcreated, err
	}

	createIndex, err := client.CreateIndex(index).BodyString(string(buf)).Do(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	if !createIndex.Acknowledged {
		// Not acknowledged
		log.Println("Index could not be created", index)
		return indexcreated, err
	} else {
		indexcreated = true
	}

	return indexcreated, err
}
