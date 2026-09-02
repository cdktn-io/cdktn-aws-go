package awseks


// Experimental.
type TfAccessPolicyAssociation_AccessScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_access_policy_association#type TfAccessPolicyAssociation#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_access_policy_association#namespaces TfAccessPolicyAssociation#namespaces}.
	// Experimental.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

