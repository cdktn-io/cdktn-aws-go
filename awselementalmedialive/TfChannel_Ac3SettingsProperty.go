package awselementalmedialive


// Experimental.
type TfChannel_Ac3SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate TfChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitstream_mode TfChannel#bitstream_mode}.
	// Experimental.
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#coding_mode TfChannel#coding_mode}.
	// Experimental.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dialnorm TfChannel#dialnorm}.
	// Experimental.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#drc_profile TfChannel#drc_profile}.
	// Experimental.
	DrcProfile *string `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#lfe_filter TfChannel#lfe_filter}.
	// Experimental.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#metadata_control TfChannel#metadata_control}.
	// Experimental.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

