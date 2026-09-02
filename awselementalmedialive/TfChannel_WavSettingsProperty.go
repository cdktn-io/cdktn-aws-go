package awselementalmedialive


// Experimental.
type TfChannel_WavSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bit_depth TfChannel#bit_depth}.
	// Experimental.
	BitDepth *float64 `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#coding_mode TfChannel#coding_mode}.
	// Experimental.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sample_rate TfChannel#sample_rate}.
	// Experimental.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

