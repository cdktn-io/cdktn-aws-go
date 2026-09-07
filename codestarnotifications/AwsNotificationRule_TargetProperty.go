package codestarnotifications


// Experimental.
type AwsNotificationRule_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarnotifications_notification_rule#address AwsNotificationRule#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarnotifications_notification_rule#type AwsNotificationRule#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

