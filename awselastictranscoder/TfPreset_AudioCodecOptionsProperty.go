package awselastictranscoder


// Experimental.
type TfPreset_AudioCodecOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#bit_depth TfPreset#bit_depth}.
	// Experimental.
	BitDepth *string `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#bit_order TfPreset#bit_order}.
	// Experimental.
	BitOrder *string `field:"optional" json:"bitOrder" yaml:"bitOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#profile TfPreset#profile}.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#signed TfPreset#signed}.
	// Experimental.
	Signed *string `field:"optional" json:"signed" yaml:"signed"`
}

