package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_WavSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bit_depth AwsMedialiveChannel#bit_depth}.
	// Experimental.
	BitDepth *float64 `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#coding_mode AwsMedialiveChannel#coding_mode}.
	// Experimental.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sample_rate AwsMedialiveChannel#sample_rate}.
	// Experimental.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

