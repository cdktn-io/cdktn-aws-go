package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object TfFlow#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#error_handling_config TfFlow#error_handling_config}
	// Experimental.
	ErrorHandlingConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigProperty `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#id_field_names TfFlow#id_field_names}.
	// Experimental.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#write_operation_type TfFlow#write_operation_type}.
	// Experimental.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

