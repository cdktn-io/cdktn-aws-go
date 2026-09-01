package awsappflow


// Experimental.
type AwsAppflowFlow_GlueDataCatalogProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#database_name AwsAppflowFlow#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#role_arn AwsAppflowFlow#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#table_prefix AwsAppflowFlow#table_prefix}.
	// Experimental.
	TablePrefix *string `field:"required" json:"tablePrefix" yaml:"tablePrefix"`
}

