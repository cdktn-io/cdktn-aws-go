package awselementalmedialive


// Experimental.
type TfChannel_Mp2SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate TfChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#coding_mode TfChannel#coding_mode}.
	// Experimental.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sample_rate TfChannel#sample_rate}.
	// Experimental.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

