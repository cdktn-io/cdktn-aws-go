package elementalmedialive


// Experimental.
type AwsChannel_InputLossBehaviorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#black_frame_msec AwsChannel#black_frame_msec}.
	// Experimental.
	BlackFrameMsec *float64 `field:"optional" json:"blackFrameMsec" yaml:"blackFrameMsec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_image_color AwsChannel#input_loss_image_color}.
	// Experimental.
	InputLossImageColor *string `field:"optional" json:"inputLossImageColor" yaml:"inputLossImageColor"`
	// input_loss_image_slate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_image_slate AwsChannel#input_loss_image_slate}
	// Experimental.
	InputLossImageSlate *AwsChannel_InputLossImageSlateProperty `field:"optional" json:"inputLossImageSlate" yaml:"inputLossImageSlate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_image_type AwsChannel#input_loss_image_type}.
	// Experimental.
	InputLossImageType *string `field:"optional" json:"inputLossImageType" yaml:"inputLossImageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#repeat_frame_msec AwsChannel#repeat_frame_msec}.
	// Experimental.
	RepeatFrameMsec *float64 `field:"optional" json:"repeatFrameMsec" yaml:"repeatFrameMsec"`
}

