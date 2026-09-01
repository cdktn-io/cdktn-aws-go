package awsamplify


// Experimental.
type AwsAmplifyDomainAssociation_SubDomainProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_domain_association#branch_name AwsAmplifyDomainAssociation#branch_name}.
	// Experimental.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_domain_association#prefix AwsAmplifyDomainAssociation#prefix}.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

