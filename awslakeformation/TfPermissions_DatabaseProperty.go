package awslakeformation


// Experimental.
type TfPermissions_DatabaseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#name TfPermissions#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#catalog_id TfPermissions#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

