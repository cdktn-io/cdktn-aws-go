package awsvpc


// Experimental.
type DataTfPrefixList_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/prefix_list#name DataTfPrefixList#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/prefix_list#values DataTfPrefixList#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

