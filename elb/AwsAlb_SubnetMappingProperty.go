package elb


// Experimental.
type AwsAlb_SubnetMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#subnet_id AwsAlb#subnet_id}.
	// Experimental.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#allocation_id AwsAlb#allocation_id}.
	// Experimental.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#ipv6_address AwsAlb#ipv6_address}.
	// Experimental.
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#private_ipv4_address AwsAlb#private_ipv4_address}.
	// Experimental.
	PrivateIpv4Address *string `field:"optional" json:"privateIpv4Address" yaml:"privateIpv4Address"`
}

