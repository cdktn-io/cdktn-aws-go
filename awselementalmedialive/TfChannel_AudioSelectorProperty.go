package awselementalmedialive


// Experimental.
type TfChannel_AudioSelectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name TfChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// selector_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#selector_settings TfChannel#selector_settings}
	// Experimental.
	SelectorSettings *TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

