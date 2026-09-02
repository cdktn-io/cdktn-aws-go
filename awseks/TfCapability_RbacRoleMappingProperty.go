package awseks


// Experimental.
type TfCapability_RbacRoleMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#role TfCapability#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#identity TfCapability#identity}
	// Experimental.
	Identity interface{} `field:"optional" json:"identity" yaml:"identity"`
}

