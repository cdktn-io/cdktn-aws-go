package vpc


// Experimental.
type DataAwsSecurityGroup_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/security_group#name DataAwsSecurityGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/security_group#values DataAwsSecurityGroup#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

