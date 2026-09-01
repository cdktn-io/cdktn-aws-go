package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty struct {
	// input_lambda_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_lambda_processor AwsKinesisanalyticsv2Application#input_lambda_processor}
	// Experimental.
	InputLambdaProcessor *AwsKinesisanalyticsv2Application_InputLambdaProcessorProperty `field:"required" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

