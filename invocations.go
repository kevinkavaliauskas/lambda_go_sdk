package lambda_go_sdk

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func invokeRenderLambda(options RemotionOptions) (*RemotionRenderResponse, error) {

	// Create a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(options.Region)},
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new Lambda client
	svc := lambda.New(sess)

	internalParams, validateError := constructRenderInternals(&options)

	log.Printf("Internal params: %v", internalParams)

	if validateError != nil {
		log.Printf("Error validating options: %v", validateError)
		return nil, validateError
	}

	internalParamJsonObject, marshallingError := json.Marshal(internalParams)
	if marshallingError != nil {

		log.Printf("Error marshalling internal params: %v", marshallingError)
		return nil, marshallingError
	}
	log.Printf("Invocation payload: %v", internalParamJsonObject)

	invocationPayload := &lambda.InvokeInput{
		FunctionName: aws.String(options.FunctionName),
		Payload:      internalParamJsonObject,
	}

	log.Printf("Invocation payload: %v", invocationPayload)

	// Invoke Lambda function
	invocationResult, invocationError := svc.Invoke(invocationPayload)

	log.Printf("Invocation error: %v", invocationResult.FunctionError)

	if invocationError != nil {
		log.Printf("Error invoking Lambda function %s: %v", options.FunctionName, invocationError)
		return nil, invocationError
	}

	// Log the raw payload and any function error
	log.Printf("Raw payload from Lambda %s: %s", options.FunctionName, string(invocationResult.Payload))
	if invocationResult.FunctionError != nil {
		log.Printf("Lambda function %s executed with error: %s. Payload: %s", options.FunctionName, *invocationResult.FunctionError, string(invocationResult.Payload))
	}

	// Unmarshal response from Lambda function
	var renderResponseOutput RemotionRenderResponse

	responseMarshallingError := json.Unmarshal(invocationResult.Payload, &renderResponseOutput)

	if responseMarshallingError != nil {
		log.Printf("Error unmarshalling response: %v", responseMarshallingError)
		log.Printf("Payload: %s", string(invocationResult.Payload))
		return nil, responseMarshallingError
	}

	return &renderResponseOutput, nil
}

func invokeRenderProgressLambda(config RenderConfig) (*RenderProgress, error) {

	// Create a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(config.Region)},
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new Lambda client
	svc := lambda.New(sess)

	internalParams, validateError := constructGetProgressInternals(&config)

	if validateError != nil {
		return nil, validateError
	}

	internalParamsJSON, marshallingError := json.Marshal(internalParams)
	if marshallingError != nil {

		return nil, marshallingError
	}

	invocationParams := &lambda.InvokeInput{
		FunctionName: aws.String(config.FunctionName),
		Payload:      internalParamsJSON,
	}

	// Invoke Lambda function
	invokeResult, invokeError := svc.Invoke(invocationParams)

	if invokeError != nil {
		log.Printf("Error invoking Lambda function %s: %v", config.FunctionName, invokeError)
		return nil, invokeError
	}

	// Log the raw payload and any function error
	log.Printf("Raw payload from Lambda %s: %s", config.FunctionName, string(invokeResult.Payload))
	if invokeResult.FunctionError != nil {
		log.Printf("Lambda function %s executed with error: %s. Payload: %s", config.FunctionName, *invokeResult.FunctionError, string(invokeResult.Payload))
	}

	// Unmarshal response from Lambda function
	var renderProgressOutput RenderProgress

	resultUnmarshallError := json.Unmarshal(invokeResult.Payload, &renderProgressOutput)
	if resultUnmarshallError != nil {
		return nil, resultUnmarshallError
	}

	return &renderProgressOutput, nil
}
