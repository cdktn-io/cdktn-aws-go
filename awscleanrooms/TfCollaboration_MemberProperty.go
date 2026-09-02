package awscleanrooms


// Experimental.
type TfCollaboration_MemberProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#account_id TfCollaboration#account_id}.
	// Experimental.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#display_name TfCollaboration#display_name}.
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#member_abilities TfCollaboration#member_abilities}.
	// Experimental.
	MemberAbilities *[]*string `field:"required" json:"memberAbilities" yaml:"memberAbilities"`
}

