package awsvpc


// Experimental.
type DataTfSecurityGroup_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/security_group#name DataTfSecurityGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/security_group#values DataTfSecurityGroup#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

