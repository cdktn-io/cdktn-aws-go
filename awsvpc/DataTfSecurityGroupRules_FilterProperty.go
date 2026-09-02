package awsvpc


// Experimental.
type DataTfSecurityGroupRules_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_security_group_rules#name DataTfSecurityGroupRules#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_security_group_rules#values DataTfSecurityGroupRules#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

