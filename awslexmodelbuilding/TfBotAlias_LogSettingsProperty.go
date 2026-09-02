package awslexmodelbuilding


// Experimental.
type TfBotAlias_LogSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#destination TfBotAlias#destination}.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#log_type TfBotAlias#log_type}.
	// Experimental.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#resource_arn TfBotAlias#resource_arn}.
	// Experimental.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot_alias#kms_key_arn TfBotAlias#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

