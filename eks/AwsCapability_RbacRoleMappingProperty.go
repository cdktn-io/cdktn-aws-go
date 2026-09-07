package eks


// Experimental.
type AwsCapability_RbacRoleMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#role AwsCapability#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#identity AwsCapability#identity}
	// Experimental.
	Identity interface{} `field:"optional" json:"identity" yaml:"identity"`
}

