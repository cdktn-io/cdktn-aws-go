package awslambda


// Experimental.
type TfFunctionEventInvokeConfig_DestinationConfigProperty struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_event_invoke_config#on_failure TfFunctionEventInvokeConfig#on_failure}
	// Experimental.
	OnFailure *TfFunctionEventInvokeConfig_OnFailureProperty `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_event_invoke_config#on_success TfFunctionEventInvokeConfig#on_success}
	// Experimental.
	OnSuccess *TfFunctionEventInvokeConfig_OnSuccessProperty `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

