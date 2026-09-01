package awsvpc


// Experimental.
type AwsEc2ManagedPrefixList_EntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_managed_prefix_list#cidr AwsEc2ManagedPrefixList#cidr}.
	// Experimental.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_managed_prefix_list#description AwsEc2ManagedPrefixList#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

