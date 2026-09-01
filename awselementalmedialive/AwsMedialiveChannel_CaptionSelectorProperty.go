package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_CaptionSelectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsMedialiveChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code AwsMedialiveChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// selector_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#selector_settings AwsMedialiveChannel#selector_settings}
	// Experimental.
	SelectorSettings *AwsMedialiveChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

