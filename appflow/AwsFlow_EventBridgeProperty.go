package appflow


// Experimental.
type AwsFlow_EventBridgeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object AwsFlow#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#error_handling_config AwsFlow#error_handling_config}
	// Experimental.
	ErrorHandlingConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfigProperty `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

