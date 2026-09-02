package awsec2


// Experimental.
type DataTfInstances_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instances#name DataTfInstances#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instances#values DataTfInstances#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

