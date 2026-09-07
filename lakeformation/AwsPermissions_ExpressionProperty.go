package lakeformation


// Experimental.
type AwsPermissions_ExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#key AwsPermissions#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#values AwsPermissions#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

