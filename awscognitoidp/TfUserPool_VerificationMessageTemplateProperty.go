package awscognitoidp


// Experimental.
type TfUserPool_VerificationMessageTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#default_email_option TfUserPool#default_email_option}.
	// Experimental.
	DefaultEmailOption *string `field:"optional" json:"defaultEmailOption" yaml:"defaultEmailOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_message TfUserPool#email_message}.
	// Experimental.
	EmailMessage *string `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_message_by_link TfUserPool#email_message_by_link}.
	// Experimental.
	EmailMessageByLink *string `field:"optional" json:"emailMessageByLink" yaml:"emailMessageByLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_subject TfUserPool#email_subject}.
	// Experimental.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_subject_by_link TfUserPool#email_subject_by_link}.
	// Experimental.
	EmailSubjectByLink *string `field:"optional" json:"emailSubjectByLink" yaml:"emailSubjectByLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sms_message TfUserPool#sms_message}.
	// Experimental.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

