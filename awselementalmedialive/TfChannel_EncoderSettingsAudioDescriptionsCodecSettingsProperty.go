package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty struct {
	// aac_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#aac_settings TfChannel#aac_settings}
	// Experimental.
	AacSettings *TfChannel_AacSettingsProperty `field:"optional" json:"aacSettings" yaml:"aacSettings"`
	// ac3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ac3_settings TfChannel#ac3_settings}
	// Experimental.
	Ac3Settings *TfChannel_Ac3SettingsProperty `field:"optional" json:"ac3Settings" yaml:"ac3Settings"`
	// eac3_atmos_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#eac3_atmos_settings TfChannel#eac3_atmos_settings}
	// Experimental.
	Eac3AtmosSettings *TfChannel_Eac3AtmosSettingsProperty `field:"optional" json:"eac3AtmosSettings" yaml:"eac3AtmosSettings"`
	// eac3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#eac3_settings TfChannel#eac3_settings}
	// Experimental.
	Eac3Settings *TfChannel_Eac3SettingsProperty `field:"optional" json:"eac3Settings" yaml:"eac3Settings"`
	// mp2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mp2_settings TfChannel#mp2_settings}
	// Experimental.
	Mp2Settings *TfChannel_Mp2SettingsProperty `field:"optional" json:"mp2Settings" yaml:"mp2Settings"`
	// pass_through_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pass_through_settings TfChannel#pass_through_settings}
	// Experimental.
	PassThroughSettings *TfChannel_PassThroughSettingsProperty `field:"optional" json:"passThroughSettings" yaml:"passThroughSettings"`
	// wav_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#wav_settings TfChannel#wav_settings}
	// Experimental.
	WavSettings *TfChannel_WavSettingsProperty `field:"optional" json:"wavSettings" yaml:"wavSettings"`
}

