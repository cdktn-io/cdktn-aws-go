package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#instance_url AwsAppflowConnectorProfile#instance_url}.
	// Experimental.
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#is_sandbox_environment AwsAppflowConnectorProfile#is_sandbox_environment}.
	// Experimental.
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#use_privatelink_for_metadata_and_authorization AwsAppflowConnectorProfile#use_privatelink_for_metadata_and_authorization}.
	// Experimental.
	UsePrivatelinkForMetadataAndAuthorization interface{} `field:"optional" json:"usePrivatelinkForMetadataAndAuthorization" yaml:"usePrivatelinkForMetadataAndAuthorization"`
}

