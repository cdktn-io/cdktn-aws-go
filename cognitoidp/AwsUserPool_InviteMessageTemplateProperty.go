package cognitoidp


// Experimental.
type AwsUserPool_InviteMessageTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_message AwsUserPool#email_message}.
	// Experimental.
	EmailMessage *string `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_subject AwsUserPool#email_subject}.
	// Experimental.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sms_message AwsUserPool#sms_message}.
	// Experimental.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

