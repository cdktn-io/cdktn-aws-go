package awsvpc


// Experimental.
type DataTfSecurityGroupRule_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_security_group_rule#name DataTfSecurityGroupRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_security_group_rule#values DataTfSecurityGroupRule#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

