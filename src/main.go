package main

import (
    "errors"
    "log"

    "github.com/tidwall/gjson"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

var (
    // ErrNameNotProvided is thrown when a name is not provided
    ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // stdout and stderr are sent to AWS CloudWatch Logs
    log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

    message := gjson.Get(request.Body, "exampleString").String()
    vowels := []byte{'a', 'e', 'i', 'o', 'u'}
    for (!contains(vowels, message[0])) {
        message = message[1:] + string(message[0])
    }
    message = message + "ay"

    return events.APIGatewayProxyResponse{
        Body:       "Translation: " + message,
	StatusCode: 200,
    }, nil
}

func main() {
    lambda.Start(Handler)
}

func contains(s []byte, e byte) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
