package awsvpcipam


// Experimental.
type DataTfPoolCidrs_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool_cidrs#name DataTfPoolCidrs#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool_cidrs#values DataTfPoolCidrs#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

