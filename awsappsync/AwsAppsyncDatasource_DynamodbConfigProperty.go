package awsappsync


// Experimental.
type AwsAppsyncDatasource_DynamodbConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#table_name AwsAppsyncDatasource#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// delta_sync_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#delta_sync_config AwsAppsyncDatasource#delta_sync_config}
	// Experimental.
	DeltaSyncConfig *AwsAppsyncDatasource_DeltaSyncConfigProperty `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#region AwsAppsyncDatasource#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#use_caller_credentials AwsAppsyncDatasource#use_caller_credentials}.
	// Experimental.
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#versioned AwsAppsyncDatasource#versioned}.
	// Experimental.
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

