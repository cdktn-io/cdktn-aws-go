package awsecs


// Experimental.
type AwsEcsDaemon_AlarmsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#alarm_names AwsEcsDaemon#alarm_names}.
	// Experimental.
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#enable AwsEcsDaemon#enable}.
	// Experimental.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

