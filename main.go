package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/danesparza/mapimage/data"
)

var (
	// BuildVersion contains the version information for the app
	BuildVersion = "Unknown"

	// CommitID is the git commitId for the app.  It's filled in as
	// part of the automated build
	CommitID string
)

// Message is a custom struct event type to handle the Lambda input
type Message struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
	Zoom int     `json:"zoom"`
}

// HandleRequest handles the AWS lambda request
func HandleRequest(ctx context.Context, msg Message) (data.MapImageResponse, error) {
	xray.Configure(xray.Config{LogLevel: "trace"})
	_, seg := xray.BeginSegment(ctx, "mapimage-lambda-handler")

	//	Call the method
	response, err := data.GetMapImageForCoordinates(msg.Lat, msg.Long, msg.Zoom)
	if err != nil {
		seg.Close(err)
		log.Fatalf("problem getting map image: %v", err)
	}

	//	Set the service version information:
	response.Version = fmt.Sprintf("%s.%s", BuildVersion, CommitID)

	//	Close the segment
	seg.Close(nil)

	//	Return our response
	return response, nil
}

func main() {
	//	Immediately forward to Lambda
	lambda.Start(HandleRequest)
}
