package awselementalmedialive


// Experimental.
type TfChannel_AudioLanguageSelectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code TfChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_selection_policy TfChannel#language_selection_policy}.
	// Experimental.
	LanguageSelectionPolicy *string `field:"optional" json:"languageSelectionPolicy" yaml:"languageSelectionPolicy"`
}

