package awscloudwatch


// Experimental.
type AwsCloudwatchCompositeAlarm_ActionsSuppressorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#alarm AwsCloudwatchCompositeAlarm#alarm}.
	// Experimental.
	Alarm *string `field:"required" json:"alarm" yaml:"alarm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#extension_period AwsCloudwatchCompositeAlarm#extension_period}.
	// Experimental.
	ExtensionPeriod *float64 `field:"required" json:"extensionPeriod" yaml:"extensionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#wait_period AwsCloudwatchCompositeAlarm#wait_period}.
	// Experimental.
	WaitPeriod *float64 `field:"required" json:"waitPeriod" yaml:"waitPeriod"`
}

