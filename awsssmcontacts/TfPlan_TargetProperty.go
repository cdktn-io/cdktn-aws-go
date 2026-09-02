package awsssmcontacts


// Experimental.
type TfPlan_TargetProperty struct {
	// channel_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#channel_target_info TfPlan#channel_target_info}
	// Experimental.
	ChannelTargetInfo *TfPlan_ChannelTargetInfoProperty `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// contact_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#contact_target_info TfPlan#contact_target_info}
	// Experimental.
	ContactTargetInfo *TfPlan_ContactTargetInfoProperty `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}

