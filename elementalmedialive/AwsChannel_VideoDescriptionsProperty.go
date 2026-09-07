package elementalmedialive


// Experimental.
type AwsChannel_VideoDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec_settings AwsChannel#codec_settings}
	// Experimental.
	CodecSettings *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#height AwsChannel#height}.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#respond_to_afd AwsChannel#respond_to_afd}.
	// Experimental.
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scaling_behavior AwsChannel#scaling_behavior}.
	// Experimental.
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sharpness AwsChannel#sharpness}.
	// Experimental.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#width AwsChannel#width}.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

