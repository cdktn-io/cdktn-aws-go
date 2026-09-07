package elementalmedialive


// Experimental.
type AwsChannel_ChannelMappingsProperty struct {
	// input_channel_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_channel_levels AwsChannel#input_channel_levels}
	// Experimental.
	InputChannelLevels interface{} `field:"required" json:"inputChannelLevels" yaml:"inputChannelLevels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_channel AwsChannel#output_channel}.
	// Experimental.
	OutputChannel *float64 `field:"required" json:"outputChannel" yaml:"outputChannel"`
}

