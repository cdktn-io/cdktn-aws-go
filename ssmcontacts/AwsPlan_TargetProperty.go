package ssmcontacts


// Experimental.
type AwsPlan_TargetProperty struct {
	// channel_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#channel_target_info AwsPlan#channel_target_info}
	// Experimental.
	ChannelTargetInfo *AwsPlan_ChannelTargetInfoProperty `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// contact_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#contact_target_info AwsPlan#contact_target_info}
	// Experimental.
	ContactTargetInfo *AwsPlan_ContactTargetInfoProperty `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}

