package lakeformation


// Experimental.
type DataAwsPermissions_ExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#key DataAwsPermissions#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions#values DataAwsPermissions#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

