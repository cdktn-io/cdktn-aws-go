package appsync


// Experimental.
type AwsDatasource_DeltaSyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#delta_sync_table_name AwsDatasource#delta_sync_table_name}.
	// Experimental.
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#base_table_ttl AwsDatasource#base_table_ttl}.
	// Experimental.
	BaseTableTtl *float64 `field:"optional" json:"baseTableTtl" yaml:"baseTableTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#delta_sync_table_ttl AwsDatasource#delta_sync_table_ttl}.
	// Experimental.
	DeltaSyncTableTtl *float64 `field:"optional" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

