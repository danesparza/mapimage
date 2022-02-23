package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
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
	ctx, seg := xray.BeginSegment(ctx, "mapimage-lambda-handler")

	//	Set the service version information:
	serviceVersion := fmt.Sprintf("%s.%s", BuildVersion, CommitID)

	//	Close the segment
	seg.Close(nil)

	//	Return our response
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", "user"),
		StatusCode: 200,
		Headers: map[string]string{
			"x-service-version": serviceVersion,
		},
	}, nil
}

func main() {
	//	Immediately forward to Lambda
	lambda.Start(HandleRequest)
}
