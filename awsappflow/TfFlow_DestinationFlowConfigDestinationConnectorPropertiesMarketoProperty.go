package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object TfFlow#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#error_handling_config TfFlow#error_handling_config}
	// Experimental.
	ErrorHandlingConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigProperty `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

