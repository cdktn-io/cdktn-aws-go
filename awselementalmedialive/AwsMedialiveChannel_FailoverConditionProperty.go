package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_FailoverConditionProperty struct {
	// failover_condition_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#failover_condition_settings AwsMedialiveChannel#failover_condition_settings}
	// Experimental.
	FailoverConditionSettings *AwsMedialiveChannel_FailoverConditionSettingsProperty `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

