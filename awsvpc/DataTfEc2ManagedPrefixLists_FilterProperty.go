package awsvpc


// Experimental.
type DataTfEc2ManagedPrefixLists_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_managed_prefix_lists#name DataTfEc2ManagedPrefixLists#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_managed_prefix_lists#values DataTfEc2ManagedPrefixLists#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

