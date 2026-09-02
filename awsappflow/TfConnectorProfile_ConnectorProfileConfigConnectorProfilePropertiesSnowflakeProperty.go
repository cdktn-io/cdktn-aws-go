package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#bucket_name TfConnectorProfile#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#stage TfConnectorProfile#stage}.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#warehouse TfConnectorProfile#warehouse}.
	// Experimental.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#account_name TfConnectorProfile#account_name}.
	// Experimental.
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#bucket_prefix TfConnectorProfile#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#private_link_service_name TfConnectorProfile#private_link_service_name}.
	// Experimental.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#region TfConnectorProfile#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

