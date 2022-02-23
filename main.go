package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
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

// HandleRequest handles the AWS lambda request
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	xray.Configure(xray.Config{LogLevel: "trace"})
	_, seg := xray.BeginSegment(ctx, "mapimage-lambda-handler")

	//	Set the service version information:
	serviceVersion := fmt.Sprintf("%s.%s", BuildVersion, CommitID)

	//	Get the lat
	lat, err := strconv.ParseFloat(request.QueryStringParameters["lat"], 64)
	if lat == 0 || err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, fmt.Errorf("lat is a required parameter")
	}

	//	Get the long
	long, err := strconv.ParseFloat(request.QueryStringParameters["long"], 64)
	if long == 0 || err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, fmt.Errorf("long is a required parameter")
	}

	//	Get the zoom level
	zoom, err := strconv.Atoi(request.QueryStringParameters["zoom"])
	if zoom == 0 || err != nil {
		//	Set a default zoom of 3
		zoom = 3
	}

	//	Call the method
	imageData, err := data.GetMapImageForCoordinates(lat, long, zoom)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	//	Close the segment
	seg.Close(nil)

	//	Return our response
	return events.APIGatewayProxyResponse{
		Body:       imageData,
		StatusCode: 200,
		Headers: map[string]string{
			"x-service-version": serviceVersion,
			"Content-type":      "image/jpeg",
		},
		IsBase64Encoded: true,
	}, nil
}

func main() {
	//	Immediately forward to Lambda
	lambda.Start(HandleRequest)
}
