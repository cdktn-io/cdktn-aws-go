package networkfirewall


// Experimental.
type AwsContainerAssociation_AttributeFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#key AwsContainerAssociation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#value AwsContainerAssociation#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

