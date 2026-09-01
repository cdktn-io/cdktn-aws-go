package awsbedrockagents


// Experimental.
type AwsBedrockagentFlow_LexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#bot_alias_arn AwsBedrockagentFlow#bot_alias_arn}.
	// Experimental.
	BotAliasArn *string `field:"required" json:"botAliasArn" yaml:"botAliasArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#locale_id AwsBedrockagentFlow#locale_id}.
	// Experimental.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

