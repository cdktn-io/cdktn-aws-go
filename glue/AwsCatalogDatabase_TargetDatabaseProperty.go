package glue


// Experimental.
type AwsCatalogDatabase_TargetDatabaseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#catalog_id AwsCatalogDatabase#catalog_id}.
	// Experimental.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#database_name AwsCatalogDatabase#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#region AwsCatalogDatabase#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

