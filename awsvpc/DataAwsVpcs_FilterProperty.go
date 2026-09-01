package awsvpc


// Experimental.
type DataAwsVpcs_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpcs#name DataAwsVpcs#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpcs#values DataAwsVpcs#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

