package awscloudwatch


// Experimental.
type TfAlarmMuteRule_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_alarm_mute_rule#duration TfAlarmMuteRule#duration}.
	// Experimental.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_alarm_mute_rule#expression TfAlarmMuteRule#expression}.
	// Experimental.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_alarm_mute_rule#timezone TfAlarmMuteRule#timezone}.
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

