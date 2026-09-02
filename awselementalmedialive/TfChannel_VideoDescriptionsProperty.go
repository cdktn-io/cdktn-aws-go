package awselementalmedialive


// Experimental.
type TfChannel_VideoDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name TfChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec_settings TfChannel#codec_settings}
	// Experimental.
	CodecSettings *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#height TfChannel#height}.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#respond_to_afd TfChannel#respond_to_afd}.
	// Experimental.
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scaling_behavior TfChannel#scaling_behavior}.
	// Experimental.
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sharpness TfChannel#sharpness}.
	// Experimental.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#width TfChannel#width}.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

