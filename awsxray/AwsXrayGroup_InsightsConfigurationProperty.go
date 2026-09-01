package awsxray


// Experimental.
type AwsXrayGroup_InsightsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_group#insights_enabled AwsXrayGroup#insights_enabled}.
	// Experimental.
	InsightsEnabled interface{} `field:"required" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_group#notifications_enabled AwsXrayGroup#notifications_enabled}.
	// Experimental.
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

