package appflow


// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object_path AwsFlow#object_path}.
	// Experimental.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#error_handling_config AwsFlow#error_handling_config}
	// Experimental.
	ErrorHandlingConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#id_field_names AwsFlow#id_field_names}.
	// Experimental.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// success_response_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#success_response_handling_config AwsFlow#success_response_handling_config}
	// Experimental.
	SuccessResponseHandlingConfig *AwsFlow_SuccessResponseHandlingConfigProperty `field:"optional" json:"successResponseHandlingConfig" yaml:"successResponseHandlingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#write_operation_type AwsFlow#write_operation_type}.
	// Experimental.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

