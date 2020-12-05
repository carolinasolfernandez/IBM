package services

import (
	"bytes"
	"encoding/json"
	"github.com/carolinasolfernandez/IBM/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//CallSubmit calls the IBM submit API
func CallSubmit(submit models.Submit) (status int) {
	log.Println(submit)
	var apiIBM = "http://172.21.86.186:5000/submit"

	body, _ := json.Marshal(submit)

	log.Println("Request: " + string(body))

	req, err := http.NewRequest("POST", apiIBM, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 5 * time.Second}

	response, err := httpClient.Do(req)
	if err != nil {
		log.Println("API call error. ", err.Error())
		return 500
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body")
	}
	log.Println(string(bodyBytes))
	return response.StatusCode
}

