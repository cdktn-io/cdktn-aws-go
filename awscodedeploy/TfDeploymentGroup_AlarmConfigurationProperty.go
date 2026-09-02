package awscodedeploy


// Experimental.
type TfDeploymentGroup_AlarmConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#alarms TfDeploymentGroup#alarms}.
	// Experimental.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#enabled TfDeploymentGroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#ignore_poll_alarm_failure TfDeploymentGroup#ignore_poll_alarm_failure}.
	// Experimental.
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

