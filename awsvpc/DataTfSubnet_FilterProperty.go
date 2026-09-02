package awsvpc


// Experimental.
type DataTfSubnet_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/subnet#name DataTfSubnet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/subnet#values DataTfSubnet#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

