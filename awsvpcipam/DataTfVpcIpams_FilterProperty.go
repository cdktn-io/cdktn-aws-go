package awsvpcipam


// Experimental.
type DataTfVpcIpams_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipams#name DataTfVpcIpams#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_ipams#values DataTfVpcIpams#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

