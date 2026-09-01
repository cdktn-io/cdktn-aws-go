package awsappflow


// Experimental.
type AwsAppflowFlow_SourceFlowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_type AwsAppflowFlow#connector_type}.
	// Experimental.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// source_connector_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#source_connector_properties AwsAppflowFlow#source_connector_properties}
	// Experimental.
	SourceConnectorProperties *AwsAppflowFlow_SourceConnectorPropertiesProperty `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#api_version AwsAppflowFlow#api_version}.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_profile_name AwsAppflowFlow#connector_profile_name}.
	// Experimental.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// incremental_pull_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#incremental_pull_config AwsAppflowFlow#incremental_pull_config}
	// Experimental.
	IncrementalPullConfig *AwsAppflowFlow_IncrementalPullConfigProperty `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

