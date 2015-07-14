package qds

import (
	"./api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const DEBUG = false

const quboleAPIVersion = "latest"
const quboleAPIRootURI = "https://api.qubole.com/api/" + quboleAPIVersion
const batchSize = 100

func buildFinalURI(uriPath *string, perPage uint, pageNumber uint) *string {
	finalURI := fmt.Sprint(quboleAPIRootURI, *uriPath, "?per_page=", perPage, "&page=", pageNumber)
	if DEBUG {
		log.Println("buildFinalURI: " + finalURI)
	}
	return &finalURI
}

func SendGetRequest(quboleToken *string, uriPath string, perPage uint, pageNumber uint) (*[]byte, error) {
	uri := buildFinalURI(&uriPath, perPage, pageNumber)
	return SendHTTPRequest(quboleToken, uri, "GET")
}

func SendHTTPRequest(quboleToken *string, uri *string, requestType string) (*[]byte, error) {

	// Checking if the tone
	if len(*quboleToken) <= 0 {
		log.Fatal("You have to provide a token for the qubole API.")
	}

	if DEBUG {
		log.Println("SendHTTPRequest: " + *uri + " - " + requestType)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(requestType, *uri, nil)

	req.Header.Add("X-AUTH-TOKEN", *quboleToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if DEBUG {
		log.Println(fmt.Sprintf("SendHTTPRequest: body=%s\nErr=%s", body, err))
	}

	return &body, err

}

func GetAllSchedules(quboleToken *string) *[](api.Schedule) {
	var schedules [](api.Schedule)

	var i uint = 1

	for {
		var data api.Scheduler
		body, err := SendGetRequest(quboleToken, "/scheduler", batchSize, i)
		err = json.Unmarshal(*body, &data)
		if err != nil {
			if DEBUG {
				log.Fatal("[ERROR] %v\n", err)
			}
		}
		schedules = append(schedules, data.Schedules...)
		i++

		if DEBUG {
			log.Printf("GetAllSchedules: NextPage=%v\nBody=%s\n", data.PagingInfo.NextPage, *body)
		}
		if data.PagingInfo.NextPage == 0 {
			break
		}
	}
	return &schedules
}
