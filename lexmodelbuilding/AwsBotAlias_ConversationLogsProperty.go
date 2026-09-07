package lexmodelbuilding


// Experimental.
type AwsBotAlias_ConversationLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#iam_role_arn AwsBotAlias#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#log_settings AwsBotAlias#log_settings}
	// Experimental.
	LogSettings interface{} `field:"optional" json:"logSettings" yaml:"logSettings"`
}

