package awscognitoidp


// Experimental.
type TfUserPool_EmailConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#configuration_set TfUserPool#configuration_set}.
	// Experimental.
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_sending_account TfUserPool#email_sending_account}.
	// Experimental.
	EmailSendingAccount *string `field:"optional" json:"emailSendingAccount" yaml:"emailSendingAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#from_email_address TfUserPool#from_email_address}.
	// Experimental.
	FromEmailAddress *string `field:"optional" json:"fromEmailAddress" yaml:"fromEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#reply_to_email_address TfUserPool#reply_to_email_address}.
	// Experimental.
	ReplyToEmailAddress *string `field:"optional" json:"replyToEmailAddress" yaml:"replyToEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#source_arn TfUserPool#source_arn}.
	// Experimental.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

