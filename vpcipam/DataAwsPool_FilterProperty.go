package vpcipam


// Experimental.
type DataAwsPool_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool#name DataAwsPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool#values DataAwsPool#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

