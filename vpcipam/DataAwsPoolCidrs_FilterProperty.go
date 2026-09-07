package vpcipam


// Experimental.
type DataAwsPoolCidrs_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool_cidrs#name DataAwsPoolCidrs#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool_cidrs#values DataAwsPoolCidrs#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

