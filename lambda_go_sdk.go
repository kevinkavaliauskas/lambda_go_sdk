package lambda_go_sdk

func RenderMediaOnLambda(input RemotionOptions) (*RemotionRenderResponse, error) {
	return invokeRenderLambda(input)
}

func GetRenderProgress(input RenderConfig) (*RenderProgress, error) {
	return invokeRenderProgressLambda(input)
}

// RenderStillOnLambda triggers a render for a still image (single frame).
// It accepts the same RemotionOptions struct as RenderMediaOnLambda, but sets
// the MediaType to "still" automatically so the caller does not need to worry
// about passing the correct value.
func RenderStillOnLambda(input RemotionOptions) (*RemotionRenderResponse, error) {
	input.MediaType = "still"
	return invokeRenderLambda(input)
}
