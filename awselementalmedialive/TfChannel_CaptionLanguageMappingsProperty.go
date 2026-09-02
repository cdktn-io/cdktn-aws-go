package awselementalmedialive


// Experimental.
type TfChannel_CaptionLanguageMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_channel TfChannel#caption_channel}.
	// Experimental.
	CaptionChannel *float64 `field:"required" json:"captionChannel" yaml:"captionChannel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code TfChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_description TfChannel#language_description}.
	// Experimental.
	LanguageDescription *string `field:"required" json:"languageDescription" yaml:"languageDescription"`
}

