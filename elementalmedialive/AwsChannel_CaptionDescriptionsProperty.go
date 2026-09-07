package elementalmedialive


// Experimental.
type AwsChannel_CaptionDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_selector_name AwsChannel#caption_selector_name}.
	// Experimental.
	CaptionSelectorName *string `field:"required" json:"captionSelectorName" yaml:"captionSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#accessibility AwsChannel#accessibility}.
	// Experimental.
	Accessibility *string `field:"optional" json:"accessibility" yaml:"accessibility"`
	// destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination_settings AwsChannel#destination_settings}
	// Experimental.
	DestinationSettings *AwsChannel_DestinationSettingsProperty `field:"optional" json:"destinationSettings" yaml:"destinationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code AwsChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_description AwsChannel#language_description}.
	// Experimental.
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

