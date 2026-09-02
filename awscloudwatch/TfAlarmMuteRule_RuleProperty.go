package awscloudwatch


// Experimental.
type TfAlarmMuteRule_RuleProperty struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_alarm_mute_rule#schedule TfAlarmMuteRule#schedule}
	// Experimental.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

