package awselementalmedialive


// Experimental.
type TfChannel_GlobalConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#initial_audio_gain TfChannel#initial_audio_gain}.
	// Experimental.
	InitialAudioGain *float64 `field:"optional" json:"initialAudioGain" yaml:"initialAudioGain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_end_action TfChannel#input_end_action}.
	// Experimental.
	InputEndAction *string `field:"optional" json:"inputEndAction" yaml:"inputEndAction"`
	// input_loss_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_behavior TfChannel#input_loss_behavior}
	// Experimental.
	InputLossBehavior *TfChannel_InputLossBehaviorProperty `field:"optional" json:"inputLossBehavior" yaml:"inputLossBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_locking_mode TfChannel#output_locking_mode}.
	// Experimental.
	OutputLockingMode *string `field:"optional" json:"outputLockingMode" yaml:"outputLockingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_timing_source TfChannel#output_timing_source}.
	// Experimental.
	OutputTimingSource *string `field:"optional" json:"outputTimingSource" yaml:"outputTimingSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#support_low_framerate_inputs TfChannel#support_low_framerate_inputs}.
	// Experimental.
	SupportLowFramerateInputs *string `field:"optional" json:"supportLowFramerateInputs" yaml:"supportLowFramerateInputs"`
}

