package awsssmcontacts


// Experimental.
type AwsSsmcontactsPlan_TargetProperty struct {
	// channel_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#channel_target_info AwsSsmcontactsPlan#channel_target_info}
	// Experimental.
	ChannelTargetInfo *AwsSsmcontactsPlan_ChannelTargetInfoProperty `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// contact_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#contact_target_info AwsSsmcontactsPlan#contact_target_info}
	// Experimental.
	ContactTargetInfo *AwsSsmcontactsPlan_ContactTargetInfoProperty `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}

