package awselementalmedialive


// Experimental.
type TfChannel_HlsOutputSettingsProperty struct {
	// hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_settings TfChannel#hls_settings}
	// Experimental.
	HlsSettings *TfChannel_HlsSettingsProperty `field:"required" json:"hlsSettings" yaml:"hlsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#h265_packaging_type TfChannel#h265_packaging_type}.
	// Experimental.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name_modifier TfChannel#name_modifier}.
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segment_modifier TfChannel#segment_modifier}.
	// Experimental.
	SegmentModifier *string `field:"optional" json:"segmentModifier" yaml:"segmentModifier"`
}

