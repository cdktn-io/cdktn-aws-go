package appsync


// Experimental.
type AwsDatasource_HttpEndpointConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#aws_secret_store_arn AwsDatasource#aws_secret_store_arn}.
	// Experimental.
	AwsSecretStoreArn *string `field:"required" json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#db_cluster_identifier AwsDatasource#db_cluster_identifier}.
	// Experimental.
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#database_name AwsDatasource#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#region AwsDatasource#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#schema AwsDatasource#schema}.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

