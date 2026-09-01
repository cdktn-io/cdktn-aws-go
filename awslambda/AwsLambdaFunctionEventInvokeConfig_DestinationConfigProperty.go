package awslambda


// Experimental.
type AwsLambdaFunctionEventInvokeConfig_DestinationConfigProperty struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_event_invoke_config#on_failure AwsLambdaFunctionEventInvokeConfig#on_failure}
	// Experimental.
	OnFailure *AwsLambdaFunctionEventInvokeConfig_OnFailureProperty `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_event_invoke_config#on_success AwsLambdaFunctionEventInvokeConfig#on_success}
	// Experimental.
	OnSuccess *AwsLambdaFunctionEventInvokeConfig_OnSuccessProperty `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

