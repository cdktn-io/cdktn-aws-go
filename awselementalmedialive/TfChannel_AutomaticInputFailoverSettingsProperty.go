package awselementalmedialive


// Experimental.
type TfChannel_AutomaticInputFailoverSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#secondary_input_id TfChannel#secondary_input_id}.
	// Experimental.
	SecondaryInputId *string `field:"required" json:"secondaryInputId" yaml:"secondaryInputId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#error_clear_time_msec TfChannel#error_clear_time_msec}.
	// Experimental.
	ErrorClearTimeMsec *float64 `field:"optional" json:"errorClearTimeMsec" yaml:"errorClearTimeMsec"`
	// failover_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#failover_condition TfChannel#failover_condition}
	// Experimental.
	FailoverCondition interface{} `field:"optional" json:"failoverCondition" yaml:"failoverCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_preference TfChannel#input_preference}.
	// Experimental.
	InputPreference *string `field:"optional" json:"inputPreference" yaml:"inputPreference"`
}

