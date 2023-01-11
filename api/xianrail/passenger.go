package xianrail

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/imroc/req/v3"
	"io"
	"log"
	"net/http"
	"time"
	"xianrail_exporter/data/passenger"
)

var httpClient = http.DefaultClient
var reqClient *req.Client

const RETRY_LIMIT = 5

func RequestAlarm() (*passenger.Alarm, error) {
	reqClient = req.C(). // Use C() to create a client and set with chainable client settings.
				SetUserAgent("my-custom-client").
				SetTimeout(5 * time.Second)
	for i := 0; i < RETRY_LIMIT; i++ {
		request, _ := sling.New().Post("https://www.xianrail.com/api/xamo-app-api/api/app/trip/psgAlarm").Request()
		response, err := httpClient.Do(request)
		if err != nil {
			return nil, err
		}
		body, _ := io.ReadAll(response.Body)
		data := passenger.Alarm{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Printf("parsing body failed, error %v, data: %s", err, string(body))
			log.Printf("retrying(%d)...", i+1)
			continue
		}
		fmt.Println(response.Status)

		reqRequest, err := reqClient.R(). // Use R() to create a request and set with chainable request settings.
							Post("https://www.xianrail.com/api/xamo-app-api/api/app/trip/psgAlarm")
		fmt.Println(reqRequest.Status)
		reqBody, _ := io.ReadAll(reqRequest.Body)
		reqData := passenger.Alarm{}
		err = json.Unmarshal(reqBody, &reqData)
		if err != nil {
			log.Printf("parsing body failed, error %v, data: %s", err, string(body))
			log.Printf("retrying(%d)...", i+1)
			continue
		}

		_ = response.Body.Close()
		_ = reqRequest.Body.Close()

		return &reqData, nil
	}
	return nil, errors.New(fmt.Sprintf("retry limit(%d) exceeded", RETRY_LIMIT))
}
