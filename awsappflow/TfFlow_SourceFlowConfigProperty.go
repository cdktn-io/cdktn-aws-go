package awsappflow


// Experimental.
type TfFlow_SourceFlowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_type TfFlow#connector_type}.
	// Experimental.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// source_connector_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#source_connector_properties TfFlow#source_connector_properties}
	// Experimental.
	SourceConnectorProperties *TfFlow_SourceConnectorPropertiesProperty `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#api_version TfFlow#api_version}.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_profile_name TfFlow#connector_profile_name}.
	// Experimental.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// incremental_pull_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#incremental_pull_config TfFlow#incremental_pull_config}
	// Experimental.
	IncrementalPullConfig *TfFlow_IncrementalPullConfigProperty `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

