package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
o		StatusCode: 200,
		Body:       "Hola desde una lambda escrita en go!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
