package awsvpcipam


// Experimental.
type DataTfPool_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool#name DataTfPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipam_pool#values DataTfPool#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

