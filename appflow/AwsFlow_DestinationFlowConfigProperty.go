package appflow


// Experimental.
type AwsFlow_DestinationFlowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_type AwsFlow#connector_type}.
	// Experimental.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// destination_connector_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#destination_connector_properties AwsFlow#destination_connector_properties}
	// Experimental.
	DestinationConnectorProperties *AwsFlow_DestinationConnectorPropertiesProperty `field:"required" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#api_version AwsFlow#api_version}.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_profile_name AwsFlow#connector_profile_name}.
	// Experimental.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
}

