package awscognitoidp


// Experimental.
type AwsCognitoUserPool_AdminCreateUserConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#allow_admin_create_user_only AwsCognitoUserPool#allow_admin_create_user_only}.
	// Experimental.
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// invite_message_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#invite_message_template AwsCognitoUserPool#invite_message_template}
	// Experimental.
	InviteMessageTemplate *AwsCognitoUserPool_InviteMessageTemplateProperty `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
}

