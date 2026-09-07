package elementalmedialive


// Experimental.
type AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty struct {
	// aac_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#aac_settings AwsChannel#aac_settings}
	// Experimental.
	AacSettings *AwsChannel_AacSettingsProperty `field:"optional" json:"aacSettings" yaml:"aacSettings"`
	// ac3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ac3_settings AwsChannel#ac3_settings}
	// Experimental.
	Ac3Settings *AwsChannel_Ac3SettingsProperty `field:"optional" json:"ac3Settings" yaml:"ac3Settings"`
	// eac3_atmos_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#eac3_atmos_settings AwsChannel#eac3_atmos_settings}
	// Experimental.
	Eac3AtmosSettings *AwsChannel_Eac3AtmosSettingsProperty `field:"optional" json:"eac3AtmosSettings" yaml:"eac3AtmosSettings"`
	// eac3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#eac3_settings AwsChannel#eac3_settings}
	// Experimental.
	Eac3Settings *AwsChannel_Eac3SettingsProperty `field:"optional" json:"eac3Settings" yaml:"eac3Settings"`
	// mp2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mp2_settings AwsChannel#mp2_settings}
	// Experimental.
	Mp2Settings *AwsChannel_Mp2SettingsProperty `field:"optional" json:"mp2Settings" yaml:"mp2Settings"`
	// pass_through_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pass_through_settings AwsChannel#pass_through_settings}
	// Experimental.
	PassThroughSettings *AwsChannel_PassThroughSettingsProperty `field:"optional" json:"passThroughSettings" yaml:"passThroughSettings"`
	// wav_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#wav_settings AwsChannel#wav_settings}
	// Experimental.
	WavSettings *AwsChannel_WavSettingsProperty `field:"optional" json:"wavSettings" yaml:"wavSettings"`
}

