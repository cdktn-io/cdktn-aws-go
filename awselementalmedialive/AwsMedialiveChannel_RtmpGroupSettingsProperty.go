package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_RtmpGroupSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ad_markers AwsMedialiveChannel#ad_markers}.
	// Experimental.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#authentication_scheme AwsMedialiveChannel#authentication_scheme}.
	// Experimental.
	AuthenticationScheme *string `field:"optional" json:"authenticationScheme" yaml:"authenticationScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#cache_full_behavior AwsMedialiveChannel#cache_full_behavior}.
	// Experimental.
	CacheFullBehavior *string `field:"optional" json:"cacheFullBehavior" yaml:"cacheFullBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#cache_length AwsMedialiveChannel#cache_length}.
	// Experimental.
	CacheLength *float64 `field:"optional" json:"cacheLength" yaml:"cacheLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_data AwsMedialiveChannel#caption_data}.
	// Experimental.
	CaptionData *string `field:"optional" json:"captionData" yaml:"captionData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_action AwsMedialiveChannel#input_loss_action}.
	// Experimental.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#restart_delay AwsMedialiveChannel#restart_delay}.
	// Experimental.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

