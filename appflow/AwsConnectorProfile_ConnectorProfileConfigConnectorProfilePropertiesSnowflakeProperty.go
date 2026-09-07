package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#bucket_name AwsConnectorProfile#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#stage AwsConnectorProfile#stage}.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#warehouse AwsConnectorProfile#warehouse}.
	// Experimental.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#account_name AwsConnectorProfile#account_name}.
	// Experimental.
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#bucket_prefix AwsConnectorProfile#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#private_link_service_name AwsConnectorProfile#private_link_service_name}.
	// Experimental.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#region AwsConnectorProfile#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

