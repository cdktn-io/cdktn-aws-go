package awslexmodelbuilding


// Experimental.
type TfBotAlias_ConversationLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#iam_role_arn TfBotAlias#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#log_settings TfBotAlias#log_settings}
	// Experimental.
	LogSettings interface{} `field:"optional" json:"logSettings" yaml:"logSettings"`
}

