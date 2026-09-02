package awsappconfig


// Experimental.
type TfEnvironment_MonitorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_environment#alarm_arn TfEnvironment#alarm_arn}.
	// Experimental.
	AlarmArn *string `field:"required" json:"alarmArn" yaml:"alarmArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_environment#alarm_role_arn TfEnvironment#alarm_role_arn}.
	// Experimental.
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

