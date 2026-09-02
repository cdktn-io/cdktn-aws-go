package awscleanrooms


// Experimental.
type TfConfiguredTable_TableReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_configured_table#database_name TfConfiguredTable#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_configured_table#table_name TfConfiguredTable#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

