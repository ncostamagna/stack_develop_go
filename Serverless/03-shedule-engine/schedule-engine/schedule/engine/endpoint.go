package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints -
type Endpoints struct {
	Schedule endpoint.Endpoint
}

// MakeEndpoints -
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Schedule: makeScheduleEndpoint(s),
	}
}

func makeScheduleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println(request)
		schedule, _ := pareseScheduleRequest(request.(ScheduleRequest))
		test, _ := s.Schedule(ctx, schedule)

		response := SuccessResponse{
			Message: "Success",
			Data:    test.list,
		}

		b, err := json.Marshal(response)
		if err != nil {
			panic("error")
		}

		return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 200,
		}, nil
	}
}

func pareseScheduleRequest(request ScheduleRequest) (*Schedule, error) {
	const (
		layoutISO  = "2006-01-02"
		layoutTime = "15:04"
	)

	startDate, _ := time.Parse(layoutISO, request.StartDate)
	endDate, _ := time.Parse(layoutISO, request.EndDate)

	var exceptions []time.Time
	for _, v := range request.Exceptions {
		exception, _ := time.Parse(layoutISO, v)
		exceptions = append(exceptions, exception)
	}

	var settings []Setting

	for _, v := range request.Settings {

		index := getSettingIndexByDay(&settings, Day(v.Day))
		if index == -1 {
			day := Day(v.Day)
			setting := &Setting{
				WeekDay:     day,
				DaySettings: nil,
			}
			settings = append(settings, *setting)
			index = len(settings) - 1
		}

		per := Periodicity(v.Periodicity)
		startTime, _ := time.Parse(layoutTime, v.StartTime)
		endTime, _ := time.Parse(layoutTime, v.EndTime)
		settings[index].AddDaySetting(startTime, endTime, per)
	}

	schedule := &Schedule{
		StartDate:  startDate,
		EndDate:    endDate,
		Settings:   settings,
		Exceptions: exceptions,
	}

	return schedule, nil
}

func getSettingIndexByDay(settings *[]Setting, Day Day) int {
	for i, setting := range *settings {
		if setting.WeekDay == Day {
			return i
		}
	}
	return -1
}
