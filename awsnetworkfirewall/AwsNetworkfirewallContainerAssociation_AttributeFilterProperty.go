package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallContainerAssociation_AttributeFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#key AwsNetworkfirewallContainerAssociation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#value AwsNetworkfirewallContainerAssociation#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

