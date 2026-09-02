package awsvpc


// Experimental.
type DataTfNetworkInterface_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/network_interface#name DataTfNetworkInterface#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/network_interface#values DataTfNetworkInterface#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

