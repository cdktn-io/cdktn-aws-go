package awslakeformation


// Experimental.
type DataTfPermissions_DataLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#arn DataTfPermissions#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#catalog_id DataTfPermissions#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

