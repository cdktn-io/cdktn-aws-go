package awsssmcontacts


// Experimental.
type TfPlan_ChannelTargetInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#contact_channel_id TfPlan#contact_channel_id}.
	// Experimental.
	ContactChannelId *string `field:"required" json:"contactChannelId" yaml:"contactChannelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#retry_interval_in_minutes TfPlan#retry_interval_in_minutes}.
	// Experimental.
	RetryIntervalInMinutes *float64 `field:"optional" json:"retryIntervalInMinutes" yaml:"retryIntervalInMinutes"`
}

