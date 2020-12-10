package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// SuccessResponse struct
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Test struct {
	Body string `json:"body"`
}

// ScheduleRequest struct
type ScheduleRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Settings  []struct {
		StartTime   string `json:"start_time"`
		EndTime     string `json:"end_time"`
		Day         int    `json:"day"`
		Periodicity int    `json:"periodicity"`
	} `json:"settings"`
	Exceptions []string `json:"exceptions"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeScheduleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("Test")
	var req ScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeAWSLambdaResponse(_ context.Context, response interface{}) ([]byte, error) {
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

func decodeAWSLambdaRequest(_ context.Context, payload []byte) (interface{}, error) {
	var req ScheduleRequest
	var test Test
	fmt.Println("test nahuel")
	err := json.Unmarshal(payload, &test)
	if err != nil {
		return req, err
	}
	err = json.Unmarshal([]byte(test.Body), &req)
	if err != nil {
		return req, err
	}
	log.Println(req)

	return req, nil
}
