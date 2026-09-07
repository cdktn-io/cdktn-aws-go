package elementalmedialive


// Experimental.
type AwsChannel_FailoverConditionProperty struct {
	// failover_condition_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#failover_condition_settings AwsChannel#failover_condition_settings}
	// Experimental.
	FailoverConditionSettings *AwsChannel_FailoverConditionSettingsProperty `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

