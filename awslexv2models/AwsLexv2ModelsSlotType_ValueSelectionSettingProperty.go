package awslexv2models


// Experimental.
type AwsLexv2ModelsSlotType_ValueSelectionSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#resolution_strategy AwsLexv2ModelsSlotType#resolution_strategy}.
	// Experimental.
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// advanced_recognition_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#advanced_recognition_setting AwsLexv2ModelsSlotType#advanced_recognition_setting}
	// Experimental.
	AdvancedRecognitionSetting interface{} `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// regex_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#regex_filter AwsLexv2ModelsSlotType#regex_filter}
	// Experimental.
	RegexFilter interface{} `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}

