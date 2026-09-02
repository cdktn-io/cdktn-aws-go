package awslakeformation


// Experimental.
type DataTfPermissions_TableProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#database_name DataTfPermissions#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#catalog_id DataTfPermissions#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#name DataTfPermissions#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#wildcard DataTfPermissions#wildcard}.
	// Experimental.
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
}

