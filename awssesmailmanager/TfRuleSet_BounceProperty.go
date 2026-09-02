package awssesmailmanager


// Experimental.
type TfRuleSet_BounceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#diagnostic_message TfRuleSet#diagnostic_message}.
	// Experimental.
	DiagnosticMessage *string `field:"required" json:"diagnosticMessage" yaml:"diagnosticMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#role_arn TfRuleSet#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#sender TfRuleSet#sender}.
	// Experimental.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#smtp_reply_code TfRuleSet#smtp_reply_code}.
	// Experimental.
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#status_code TfRuleSet#status_code}.
	// Experimental.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy TfRuleSet#action_failure_policy}.
	// Experimental.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#message TfRuleSet#message}.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

