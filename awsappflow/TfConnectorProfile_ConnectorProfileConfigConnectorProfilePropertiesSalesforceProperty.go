package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#instance_url TfConnectorProfile#instance_url}.
	// Experimental.
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#is_sandbox_environment TfConnectorProfile#is_sandbox_environment}.
	// Experimental.
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#use_privatelink_for_metadata_and_authorization TfConnectorProfile#use_privatelink_for_metadata_and_authorization}.
	// Experimental.
	UsePrivatelinkForMetadataAndAuthorization interface{} `field:"optional" json:"usePrivatelinkForMetadataAndAuthorization" yaml:"usePrivatelinkForMetadataAndAuthorization"`
}

