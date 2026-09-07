package codedeploy


// Experimental.
type AwsDeploymentGroup_AlarmConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#alarms AwsDeploymentGroup#alarms}.
	// Experimental.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#enabled AwsDeploymentGroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#ignore_poll_alarm_failure AwsDeploymentGroup#ignore_poll_alarm_failure}.
	// Experimental.
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

