package awskinesisanalyticsv2


// Experimental.
type TfApplication_InputProcessingConfigurationProperty struct {
	// input_lambda_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_lambda_processor TfApplication#input_lambda_processor}
	// Experimental.
	InputLambdaProcessor *TfApplication_InputLambdaProcessorProperty `field:"required" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

