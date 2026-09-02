package awselementalmedialive


// Experimental.
type TfChannel_FailoverConditionProperty struct {
	// failover_condition_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#failover_condition_settings TfChannel#failover_condition_settings}
	// Experimental.
	FailoverConditionSettings *TfChannel_FailoverConditionSettingsProperty `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

