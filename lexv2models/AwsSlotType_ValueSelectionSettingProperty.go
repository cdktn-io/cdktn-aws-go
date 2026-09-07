package lexv2models


// Experimental.
type AwsSlotType_ValueSelectionSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#resolution_strategy AwsSlotType#resolution_strategy}.
	// Experimental.
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// advanced_recognition_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#advanced_recognition_setting AwsSlotType#advanced_recognition_setting}
	// Experimental.
	AdvancedRecognitionSetting interface{} `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// regex_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#regex_filter AwsSlotType#regex_filter}
	// Experimental.
	RegexFilter interface{} `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}

