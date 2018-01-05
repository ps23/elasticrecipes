package elasticrecipes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jpillora/backoff"
)

func Check(elasticserver string, attempts int) bool {

	var boolstat bool

	b := &backoff.Backoff{
		//These are the defaults
		// Min:    500 * time.Millisecond,
		// Max:    100 * time.Second,
		Factor: 2,
		Jitter: false,
	}

	var i int

	for {
		i++
		log.Println("Check ES status -->", i, "times", "attempts", attempts)

		if i > attempts {
			log.Println("can't connect to ES", elasticserver, ". Giving up after ", attempts, "attempts")
			b.Reset()
			break
		}

		urlstr := elasticserver + "/_cluster/health?wait_for_status=yellow&timeout=5s"

		resp, err := http.Get(urlstr)
		if err != nil {

			d := b.Duration()
			log.Println("reconnecting ES in", d)
			time.Sleep(d)
			continue
		}
		defer resp.Body.Close()
		var jsonres map[string]interface{}

		err = json.NewDecoder(resp.Body).Decode(&jsonres)
		if err != nil {

			log.Println(err.Error())
		} else {

			status := jsonres["status"].(string)

			if status == "yellow" || status == "green" {
				log.Println("OK ES status", status)
				b.Reset()

				boolstat = true

				break
			} else {

				log.Println("ES status", status)
				d := b.Duration()
				log.Println("reconnecting ES in", d)
				time.Sleep(d)
				// continue
			}
		}
	}
	
	return boolstat
}
