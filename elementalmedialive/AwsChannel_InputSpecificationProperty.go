package elementalmedialive


// Experimental.
type AwsChannel_InputSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec AwsChannel#codec}.
	// Experimental.
	Codec *string `field:"required" json:"codec" yaml:"codec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_resolution AwsChannel#input_resolution}.
	// Experimental.
	InputResolution *string `field:"required" json:"inputResolution" yaml:"inputResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#maximum_bitrate AwsChannel#maximum_bitrate}.
	// Experimental.
	MaximumBitrate *string `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
}

