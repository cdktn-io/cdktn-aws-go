package awslakeformation


// Experimental.
type DataTfPermissions_ExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#key DataTfPermissions#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#values DataTfPermissions#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

