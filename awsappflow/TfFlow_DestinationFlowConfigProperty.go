package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_type TfFlow#connector_type}.
	// Experimental.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// destination_connector_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#destination_connector_properties TfFlow#destination_connector_properties}
	// Experimental.
	DestinationConnectorProperties *TfFlow_DestinationConnectorPropertiesProperty `field:"required" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#api_version TfFlow#api_version}.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_profile_name TfFlow#connector_profile_name}.
	// Experimental.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
}

