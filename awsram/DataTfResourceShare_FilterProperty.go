package awsram


// Experimental.
type DataTfResourceShare_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ram_resource_share#name DataTfResourceShare#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ram_resource_share#values DataTfResourceShare#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

