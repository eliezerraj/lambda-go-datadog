package handler

import(
	"github.com/rs/zerolog/log"
	"net/http"
//	"encoding/json"

	"github.com/lambda-go-datadog/internal/erro"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"

)

var childLogger = log.With().Str("handler", "DatadogHandler").Logger()

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type MessageBody struct {
	Msg *string `json:"message,omitempty"`
}

type DatadogHandler struct {
}

func (h *DatadogHandler) UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return ApiHandlerResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(erro.ErrMethodNotAllowed.Error())})
}

func NewDatadogHandler() *DatadogHandler{
	childLogger.Debug().Msg("DatadogHandler")
	return &DatadogHandler{}
}

func (h *DatadogHandler) GetVersion(version string) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetVersion")

	response := MessageBody { Msg: &version }
	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	return handlerResponse, nil
}

