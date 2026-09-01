package awsebs


// Experimental.
type DataAwsEbsVolumes_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_volumes#name DataAwsEbsVolumes#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_volumes#values DataAwsEbsVolumes#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

