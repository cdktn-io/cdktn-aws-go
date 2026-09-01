package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_CaptionDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_selector_name AwsMedialiveChannel#caption_selector_name}.
	// Experimental.
	CaptionSelectorName *string `field:"required" json:"captionSelectorName" yaml:"captionSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsMedialiveChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#accessibility AwsMedialiveChannel#accessibility}.
	// Experimental.
	Accessibility *string `field:"optional" json:"accessibility" yaml:"accessibility"`
	// destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination_settings AwsMedialiveChannel#destination_settings}
	// Experimental.
	DestinationSettings *AwsMedialiveChannel_DestinationSettingsProperty `field:"optional" json:"destinationSettings" yaml:"destinationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code AwsMedialiveChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_description AwsMedialiveChannel#language_description}.
	// Experimental.
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

