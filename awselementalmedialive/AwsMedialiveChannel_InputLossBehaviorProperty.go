package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_InputLossBehaviorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#black_frame_msec AwsMedialiveChannel#black_frame_msec}.
	// Experimental.
	BlackFrameMsec *float64 `field:"optional" json:"blackFrameMsec" yaml:"blackFrameMsec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_image_color AwsMedialiveChannel#input_loss_image_color}.
	// Experimental.
	InputLossImageColor *string `field:"optional" json:"inputLossImageColor" yaml:"inputLossImageColor"`
	// input_loss_image_slate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_image_slate AwsMedialiveChannel#input_loss_image_slate}
	// Experimental.
	InputLossImageSlate *AwsMedialiveChannel_InputLossImageSlateProperty `field:"optional" json:"inputLossImageSlate" yaml:"inputLossImageSlate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_image_type AwsMedialiveChannel#input_loss_image_type}.
	// Experimental.
	InputLossImageType *string `field:"optional" json:"inputLossImageType" yaml:"inputLossImageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#repeat_frame_msec AwsMedialiveChannel#repeat_frame_msec}.
	// Experimental.
	RepeatFrameMsec *float64 `field:"optional" json:"repeatFrameMsec" yaml:"repeatFrameMsec"`
}

