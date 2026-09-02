package awselementalmedialive


// Experimental.
type TfChannel_AudioNormalizationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#algorithm TfChannel#algorithm}.
	// Experimental.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#algorithm_control TfChannel#algorithm_control}.
	// Experimental.
	AlgorithmControl *string `field:"optional" json:"algorithmControl" yaml:"algorithmControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#target_lkfs TfChannel#target_lkfs}.
	// Experimental.
	TargetLkfs *float64 `field:"optional" json:"targetLkfs" yaml:"targetLkfs"`
}

