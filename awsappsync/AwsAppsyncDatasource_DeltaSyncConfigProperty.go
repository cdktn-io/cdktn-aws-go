package awsappsync


// Experimental.
type AwsAppsyncDatasource_DeltaSyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#delta_sync_table_name AwsAppsyncDatasource#delta_sync_table_name}.
	// Experimental.
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#base_table_ttl AwsAppsyncDatasource#base_table_ttl}.
	// Experimental.
	BaseTableTtl *float64 `field:"optional" json:"baseTableTtl" yaml:"baseTableTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#delta_sync_table_ttl AwsAppsyncDatasource#delta_sync_table_ttl}.
	// Experimental.
	DeltaSyncTableTtl *float64 `field:"optional" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

