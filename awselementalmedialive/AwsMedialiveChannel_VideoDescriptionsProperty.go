package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_VideoDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsMedialiveChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec_settings AwsMedialiveChannel#codec_settings}
	// Experimental.
	CodecSettings *AwsMedialiveChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#height AwsMedialiveChannel#height}.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#respond_to_afd AwsMedialiveChannel#respond_to_afd}.
	// Experimental.
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scaling_behavior AwsMedialiveChannel#scaling_behavior}.
	// Experimental.
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sharpness AwsMedialiveChannel#sharpness}.
	// Experimental.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#width AwsMedialiveChannel#width}.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

