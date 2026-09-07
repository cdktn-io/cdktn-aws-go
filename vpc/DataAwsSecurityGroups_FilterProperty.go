package vpc


// Experimental.
type DataAwsSecurityGroups_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/security_groups#name DataAwsSecurityGroups#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/security_groups#values DataAwsSecurityGroups#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

