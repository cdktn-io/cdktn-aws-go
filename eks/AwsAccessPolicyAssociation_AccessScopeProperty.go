package eks


// Experimental.
type AwsAccessPolicyAssociation_AccessScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_access_policy_association#type AwsAccessPolicyAssociation#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_access_policy_association#namespaces AwsAccessPolicyAssociation#namespaces}.
	// Experimental.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

