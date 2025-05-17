package lambda_go_sdk

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func constructRenderInternals(options *RemotionOptions) (*renderInternalOptions, error) {
	log.Println("constructRenderInternals PIPIPIPOOPOPOPOPOP", options.MediaType)

	inputProps, serializeError := serializeInputProps(options.InputProps, options.Region, options.MediaType, options.ForceBucketName)

	if serializeError != nil {
		log.Fatal("Error in serializing input props", serializeError)
	}
	validate := validator.New()
	validationErrors := validate.Struct(options)
	if validationErrors != nil {

		return nil, validationErrors
	}

	jpegQuality := 80
	if options.JpegQuality != 0 {
		jpegQuality = options.JpegQuality
	}

	internalParams := renderInternalOptions{
		ServeUrl:                       options.ServeUrl,
		InputProps:                     inputProps,
		Composition:                    options.Composition,
		Version:                        VERSION,
		FrameRange:                     options.FrameRange,
		OutName:                        options.OutName,
		AudioBitrate:                   options.AudioBitrate,
		VideoBitrate:                   options.VideoBitrate,
		Webhook:                        options.Webhook,
		ForceHeight:                    options.ForceHeight,
		OffthreadVideoCacheSizeInBytes: options.OffthreadVideoCacheSizeInBytes,
		OffthreadVideoThreads:          options.OffthreadVideoThreads,
		X264Preset:                     options.X264Preset,
		ForceWidth:                     options.ForceWidth,
		ApiKey:                         options.ApiKey,
		BucketName:                     options.BucketName,
		AudioCodec:                     options.AudioCodec,
		ForceBucketName:                options.ForceBucketName,
		RendererFunctionName:           &options.RendererFunctionName,
		DeleteAfter:                    options.DeleteAfter,
		Type:                           "start",
		JpegQuality:                    jpegQuality,
	}

	internalParams.Muted = options.Muted
	internalParams.PreferLossless = options.PreferLossless
	internalParams.Overwrite = options.Overwrite

	if options.RendererFunctionName == "" {
		internalParams.RendererFunctionName = nil
	}
	if options.Codec == "" {
		internalParams.Codec = "h264"
	} else {
		internalParams.Codec = options.Codec
	}
	if options.EveryNthFrame == 0 {
		internalParams.EveryNthFrame = 1
	} else {
		internalParams.EveryNthFrame = options.EveryNthFrame
	}

	if options.ImageFormat == "" {
		internalParams.ImageFormat = "jpeg"
	} else {
		internalParams.ImageFormat = options.ImageFormat
	}
	if options.Crf == 0 {
		internalParams.Crf = nil
	} else {
		internalParams.Crf = options.Crf
	}
	if options.Privacy == "" {
		internalParams.Privacy = "public"
	} else {
		internalParams.Privacy = options.Privacy
	}
	if options.ColorSpace == "" {
		internalParams.ColorSpace = nil
	} else {
		internalParams.ColorSpace = options.ColorSpace
	}
	if options.LogLevel == "" {
		internalParams.LogLevel = "info"
	} else {
		internalParams.LogLevel = options.LogLevel
	}

	if options.Scale == 0 {
		internalParams.Scale = 1
	} else {
		internalParams.Scale = options.Scale
	}

	if options.Codec == "" {
		internalParams.Codec = "h264"
	} else {
		internalParams.Codec = options.Codec
	}

	if options.MaxRetries == 0 {
		internalParams.MaxRetries = 1
	} else {
		internalParams.MaxRetries = options.MaxRetries
	}

	if options.Scale == 0 {
		internalParams.Scale = 1
	} else {
		internalParams.Scale = options.Scale
	}

	if options.ConcurrencyPerLambda == 0 {
		internalParams.ConcurrencyPerLambda = 1
	} else {
		internalParams.ConcurrencyPerLambda = options.ConcurrencyPerLambda
	}

	if options.TimeoutInMilliseconds == 0 {
		internalParams.TimeoutInMilliseconds = 30000
	} else {
		internalParams.TimeoutInMilliseconds = options.TimeoutInMilliseconds
	}
	internalParams.NumberOfGifLoops = options.NumberOfGifLoops

	if options.DownloadBehavior == nil {
		internalParams.DownloadBehavior = map[string]interface{}{
			"type": "play-in-browser",
		}
	} else {
		internalParams.DownloadBehavior = options.DownloadBehavior
	}
	if options.ChromiumOptions == nil {
		internalParams.ChromiumOptions = map[string]interface{}{}
	} else {
		internalParams.ChromiumOptions = options.ChromiumOptions
	}
	if options.EnvVariables == nil {
		internalParams.EnvVariables = map[string]interface{}{}
	} else {
		internalParams.EnvVariables = options.EnvVariables
	}
	if options.Metadata == nil {
		internalParams.Metadata = map[string]interface{}{}
	} else {
		internalParams.Metadata = options.Metadata
	}

	return &internalParams, nil
}

func constructGetProgressInternals(options *RenderConfig) (*renderProgressInternalConfig, error) {

	validate := validator.New()
	validationErrors := validate.Struct(options)
	if validationErrors != nil {

		return nil, validationErrors
	}

	logLevel := "info"
	if options.LogLevel != "" {
		logLevel = options.LogLevel
	}

	internalParams := renderProgressInternalConfig{
		RenderId:   options.RenderId,
		BucketName: options.BucketName,
		LogLevel:   logLevel,
		Type:       "status",
		Version:    VERSION,
	}

	return &internalParams, nil
}

func constructStillInternals(options *RemotionStillOptions) (*renderStillInternalOptions, error) {
	validate := validator.New()
	validationErrors := validate.Struct(options)
	if validationErrors != nil {
		return nil, validationErrors
	}

	// Serialize input props
	inputProps, serializeError := serializeInputProps(options.InputProps, options.Region, "still", options.ForceBucketName)
	if serializeError != nil {
		log.Fatal("Error in serializing input props", serializeError)
	}

	// Defaults
	jpegQuality := 80
	if options.JpegQuality != 0 {
		jpegQuality = options.JpegQuality
	}

	maxRetries := 1
	if options.MaxRetries != 0 {
		maxRetries = options.MaxRetries
	}

	logLevel := "info"
	if options.LogLevel != "" {
		logLevel = options.LogLevel
	}

	timeout := 30000
	if options.TimeoutInMilliseconds != 0 {
		timeout = options.TimeoutInMilliseconds
	}

	scale := 1.0
	if options.Scale != 0 {
		scale = options.Scale
	}

	attempt := 1
	if options.Attempt != 0 {
		attempt = options.Attempt
	}

	downloadBehavior := options.DownloadBehavior
	if downloadBehavior == nil {
		downloadBehavior = map[string]interface{}{"type": "play-in-browser"}
	}

	chromiumOptions := options.ChromiumOptions
	if chromiumOptions == nil {
		chromiumOptions = map[string]interface{}{}
	}

	envVariables := options.EnvVariables
	if envVariables == nil {
		envVariables = map[string]interface{}{}
	}

	internal := renderStillInternalOptions{
		Type:                           "still",
		Composition:                    options.Composition,
		ServeUrl:                       options.ServeUrl,
		InputProps:                     inputProps,
		ImageFormat:                    options.ImageFormat,
		Privacy:                        options.Privacy,
		Version:                        VERSION,
		TimeoutInMilliseconds:          timeout,
		MaxRetries:                     maxRetries,
		EnvVariables:                   envVariables,
		JpegQuality:                    jpegQuality,
		StorageClass:                   options.StorageClass,
		Frame:                          options.Frame,
		LogLevel:                       logLevel,
		OutName:                        options.OutName,
		ChromiumOptions:                chromiumOptions,
		Scale:                          scale,
		DownloadBehavior:               downloadBehavior,
		ForceWidth:                     options.ForceWidth,
		ApiKey:                         options.ApiKey,
		ForceHeight:                    options.ForceHeight,
		ForceBucketName:                options.ForceBucketName,
		DeleteAfter:                    options.DeleteAfter,
		Attempt:                        attempt,
		OffthreadVideoCacheSizeInBytes: options.OffthreadVideoCacheSizeInBytes,
		OffthreadVideoThreads:          options.OffthreadVideoThreads,
		Streamed:                       options.Streamed,
		ForcePathStyle:                 options.ForcePathStyle,
	}

	// Default imageFormat
	if internal.ImageFormat == "" {
		internal.ImageFormat = "jpeg"
	}

	if internal.Privacy == "" {
		internal.Privacy = "public"
	}

	return &internal, nil
}
