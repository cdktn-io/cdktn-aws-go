package awsvpc


// Experimental.
type DataTfEc2ManagedPrefixList_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_managed_prefix_list#name DataTfEc2ManagedPrefixList#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_managed_prefix_list#values DataTfEc2ManagedPrefixList#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

