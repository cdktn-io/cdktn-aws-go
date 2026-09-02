package awsssoadmin


// Experimental.
type TfPermissionSetInlinePolicy_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permission_set_inline_policy#create TfPermissionSetInlinePolicy#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permission_set_inline_policy#delete TfPermissionSetInlinePolicy#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

