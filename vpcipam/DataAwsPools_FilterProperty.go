package vpcipam


// Experimental.
type DataAwsPools_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pools#name DataAwsPools#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pools#values DataAwsPools#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

