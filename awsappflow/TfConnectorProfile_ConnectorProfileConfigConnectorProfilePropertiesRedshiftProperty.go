package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#bucket_name TfConnectorProfile#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#role_arn TfConnectorProfile#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#bucket_prefix TfConnectorProfile#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#cluster_identifier TfConnectorProfile#cluster_identifier}.
	// Experimental.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#data_api_role_arn TfConnectorProfile#data_api_role_arn}.
	// Experimental.
	DataApiRoleArn *string `field:"optional" json:"dataApiRoleArn" yaml:"dataApiRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#database_name TfConnectorProfile#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#database_url TfConnectorProfile#database_url}.
	// Experimental.
	DatabaseUrl *string `field:"optional" json:"databaseUrl" yaml:"databaseUrl"`
}

