package elementalmedialive


// Experimental.
type AwsChannel_OutputsProperty struct {
	// output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_settings AwsChannel#output_settings}
	// Experimental.
	OutputSettings *AwsChannel_OutputSettingsProperty `field:"required" json:"outputSettings" yaml:"outputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_description_names AwsChannel#audio_description_names}.
	// Experimental.
	AudioDescriptionNames *[]*string `field:"optional" json:"audioDescriptionNames" yaml:"audioDescriptionNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_description_names AwsChannel#caption_description_names}.
	// Experimental.
	CaptionDescriptionNames *[]*string `field:"optional" json:"captionDescriptionNames" yaml:"captionDescriptionNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_name AwsChannel#output_name}.
	// Experimental.
	OutputName *string `field:"optional" json:"outputName" yaml:"outputName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_description_name AwsChannel#video_description_name}.
	// Experimental.
	VideoDescriptionName *string `field:"optional" json:"videoDescriptionName" yaml:"videoDescriptionName"`
}

