package main

import(
//	"fmt"
	"os"
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lambda-go-datadog/internal/adapter/handler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

var (
	logLevel 		= zerolog.DebugLevel // InfoLevel DebugLevel
	version 		= "lambda go-datadog v 1.0"
	datadogHandler 	*handler.DatadogHandler
	response		*events.APIGatewayProxyResponse
)

func getEnv() {
	if os.Getenv("LOG_LEVEL") !=  "" {
		if (os.Getenv("LOG_LEVEL") == "DEBUG"){
			logLevel = zerolog.DebugLevel
		}else if (os.Getenv("LOG_LEVEL") == "INFO"){
			logLevel = zerolog.InfoLevel
		}else if (os.Getenv("LOG_LEVEL") == "ERROR"){
				logLevel = zerolog.ErrorLevel
		}else {
			logLevel = zerolog.InfoLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func init(){
	log.Debug().Msg("*** init")
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main(){
	log.Debug().Msg("*** lambda-go-datadog")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	datadogHandler		= handler.NewDatadogHandler()
	lambda.Start(ddlambda.WrapFunction(lambdaHandler, nil))
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("lambdaHandler")
	log.Debug().Msg("-------------------")
	log.Debug().Str("req.Body", req.Body).
				Msg("APIGateway Request.Body")
	log.Debug().Msg("--------------------")

	req, _ := http.NewRequestWithContext(ctx, "GET", "https://www.datadoghq.com", nil)
	client := http.Client{}
	client = *httptrace.WrapClient(&client)
	client.Do(req)

	ddlambda.Metric(
		"coffee_house.order_value", // Metric name
		12.45, // Metric value
		"product:latte", "order:online", // Associated tags  
	)
	s, _ := tracer.StartSpanFromContext(ctx, "child.span")
	s.Finish()
	
	switch req.HTTPMethod {
		case "GET":
			if (req.Resource == "/datadog/get"){
				//response, _ = datadogHandler.getData()
			}else if (req.Resource == "/version"){
				response, _ = datadogHandler.GetVersion(version)
			}else {
				response, _ = datadogHandler.UnhandledMethod()
			}
		case "POST":
			response, _ = datadogHandler.UnhandledMethod()
		case "DELETE":
			response, _ = datadogHandler.UnhandledMethod()
		case "PUT":
			response, _ = datadogHandler.UnhandledMethod()
		default:
			response, _ = datadogHandler.UnhandledMethod()
	}

	return response, nil
}