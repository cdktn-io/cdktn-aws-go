package awsappsync


// Experimental.
type TfDatasource_HttpEndpointConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#aws_secret_store_arn TfDatasource#aws_secret_store_arn}.
	// Experimental.
	AwsSecretStoreArn *string `field:"required" json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#db_cluster_identifier TfDatasource#db_cluster_identifier}.
	// Experimental.
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#database_name TfDatasource#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#region TfDatasource#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#schema TfDatasource#schema}.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

