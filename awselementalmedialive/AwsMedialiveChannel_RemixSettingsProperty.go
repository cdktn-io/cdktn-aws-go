package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_RemixSettingsProperty struct {
	// channel_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#channel_mappings AwsMedialiveChannel#channel_mappings}
	// Experimental.
	ChannelMappings interface{} `field:"required" json:"channelMappings" yaml:"channelMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#channels_in AwsMedialiveChannel#channels_in}.
	// Experimental.
	ChannelsIn *float64 `field:"optional" json:"channelsIn" yaml:"channelsIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#channels_out AwsMedialiveChannel#channels_out}.
	// Experimental.
	ChannelsOut *float64 `field:"optional" json:"channelsOut" yaml:"channelsOut"`
}

