package elementalmedialive


// Experimental.
type AwsChannel_Ac3SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate AwsChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitstream_mode AwsChannel#bitstream_mode}.
	// Experimental.
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#coding_mode AwsChannel#coding_mode}.
	// Experimental.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dialnorm AwsChannel#dialnorm}.
	// Experimental.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#drc_profile AwsChannel#drc_profile}.
	// Experimental.
	DrcProfile *string `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#lfe_filter AwsChannel#lfe_filter}.
	// Experimental.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#metadata_control AwsChannel#metadata_control}.
	// Experimental.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

