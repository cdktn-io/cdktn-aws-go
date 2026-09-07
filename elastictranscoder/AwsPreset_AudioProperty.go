package elastictranscoder


// Experimental.
type AwsPreset_AudioProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#audio_packing_mode AwsPreset#audio_packing_mode}.
	// Experimental.
	AudioPackingMode *string `field:"optional" json:"audioPackingMode" yaml:"audioPackingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#bit_rate AwsPreset#bit_rate}.
	// Experimental.
	BitRate *string `field:"optional" json:"bitRate" yaml:"bitRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#channels AwsPreset#channels}.
	// Experimental.
	Channels *string `field:"optional" json:"channels" yaml:"channels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#codec AwsPreset#codec}.
	// Experimental.
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#sample_rate AwsPreset#sample_rate}.
	// Experimental.
	SampleRate *string `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

