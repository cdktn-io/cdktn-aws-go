package cloudwatch


// Experimental.
type AwsAlarmMuteRule_MuteTargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_alarm_mute_rule#alarm_names AwsAlarmMuteRule#alarm_names}.
	// Experimental.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
}

