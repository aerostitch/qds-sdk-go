package qds_sdk

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

const qubole_api_version = "latest"
const qubole_api_root_uri = "https://api.qubole.com/api/" + qubole_api_version
const batch_size = 100

func buildFinalUri(uri_path *string, per_page uint, page_number uint) *string {
	final_uri := fmt.Sprint(qubole_api_root_uri, *uri_path, "?per_page=", per_page, "&page=", page_number)
	if DEBUG {
		log.Println("buildFinalUri: " + final_uri)
	}
	return &final_uri
}

func SendGetRequest(qubole_token *string, uri_path string, per_page uint, page_number uint) (*[]byte, error) {
	uri := buildFinalUri(&uri_path, per_page, page_number)
	return SendHttpRequest(qubole_token, uri, "GET")
}

func SendHttpRequest(qubole_token *string, uri *string, request_type string) (*[]byte, error) {

	// Checking if the tone
	if len(*qubole_token) <= 0 {
		log.Fatal("You have to provide a token for the qubole API.")
	}

	if DEBUG {
		log.Println("SendHttpRequest: " + *uri + " - " + request_type)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(request_type, *uri, nil)

	req.Header.Add("X-AUTH-TOKEN", *qubole_token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if DEBUG {
		log.Println(fmt.Sprintf("SendHttpRequest: body=%s\nErr=%s", body, err))
	}

	return &body, err

}

func GetAllSchedules(qubole_token *string) *[](api.Schedule) {
	var schedules [](api.Schedule)

	var i uint = 1

	for {
		var data api.Scheduler
		body, err := SendGetRequest(qubole_token, "/scheduler", batch_size, i)
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
