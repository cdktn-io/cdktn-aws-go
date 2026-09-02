package awscognitoidp


// Experimental.
type TfUserPool_AdminCreateUserConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#allow_admin_create_user_only TfUserPool#allow_admin_create_user_only}.
	// Experimental.
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// invite_message_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#invite_message_template TfUserPool#invite_message_template}
	// Experimental.
	InviteMessageTemplate *TfUserPool_InviteMessageTemplateProperty `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
}

