package awsnetworkfirewall


// Experimental.
type TfContainerAssociation_AttributeFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#key TfContainerAssociation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#value TfContainerAssociation#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

