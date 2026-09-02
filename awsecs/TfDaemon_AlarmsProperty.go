package awsecs


// Experimental.
type TfDaemon_AlarmsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#alarm_names TfDaemon#alarm_names}.
	// Experimental.
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#enable TfDaemon#enable}.
	// Experimental.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

