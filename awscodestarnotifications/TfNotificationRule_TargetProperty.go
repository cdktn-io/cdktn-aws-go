package awscodestarnotifications


// Experimental.
type TfNotificationRule_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarnotifications_notification_rule#address TfNotificationRule#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarnotifications_notification_rule#type TfNotificationRule#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

