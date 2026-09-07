package elementalmedialive


// Experimental.
type AwsChannel_InputSettingsProperty struct {
	// audio_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_selector AwsChannel#audio_selector}
	// Experimental.
	AudioSelector interface{} `field:"optional" json:"audioSelector" yaml:"audioSelector"`
	// caption_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_selector AwsChannel#caption_selector}
	// Experimental.
	CaptionSelector interface{} `field:"optional" json:"captionSelector" yaml:"captionSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#deblock_filter AwsChannel#deblock_filter}.
	// Experimental.
	DeblockFilter *string `field:"optional" json:"deblockFilter" yaml:"deblockFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#denoise_filter AwsChannel#denoise_filter}.
	// Experimental.
	DenoiseFilter *string `field:"optional" json:"denoiseFilter" yaml:"denoiseFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filter_strength AwsChannel#filter_strength}.
	// Experimental.
	FilterStrength *float64 `field:"optional" json:"filterStrength" yaml:"filterStrength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_filter AwsChannel#input_filter}.
	// Experimental.
	InputFilter *string `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// network_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#network_input_settings AwsChannel#network_input_settings}
	// Experimental.
	NetworkInputSettings *AwsChannel_NetworkInputSettingsProperty `field:"optional" json:"networkInputSettings" yaml:"networkInputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte35_pid AwsChannel#scte35_pid}.
	// Experimental.
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#smpte2038_data_preference AwsChannel#smpte2038_data_preference}.
	// Experimental.
	Smpte2038DataPreference *string `field:"optional" json:"smpte2038DataPreference" yaml:"smpte2038DataPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#source_end_behavior AwsChannel#source_end_behavior}.
	// Experimental.
	SourceEndBehavior *string `field:"optional" json:"sourceEndBehavior" yaml:"sourceEndBehavior"`
	// video_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_selector AwsChannel#video_selector}
	// Experimental.
	VideoSelector *AwsChannel_VideoSelectorProperty `field:"optional" json:"videoSelector" yaml:"videoSelector"`
}

