package kinesisanalyticsv2


// Experimental.
type AwsApplication_InputProcessingConfigurationProperty struct {
	// input_lambda_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_lambda_processor AwsApplication#input_lambda_processor}
	// Experimental.
	InputLambdaProcessor *AwsApplication_InputLambdaProcessorProperty `field:"required" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

