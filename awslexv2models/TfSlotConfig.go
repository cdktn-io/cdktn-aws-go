package awslexv2models

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#bot_id TfSlot#bot_id}.
	// Experimental.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#bot_version TfSlot#bot_version}.
	// Experimental.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#intent_id TfSlot#intent_id}.
	// Experimental.
	IntentId *string `field:"required" json:"intentId" yaml:"intentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#locale_id TfSlot#locale_id}.
	// Experimental.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#name TfSlot#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#description TfSlot#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// multiple_values_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#multiple_values_setting TfSlot#multiple_values_setting}
	// Experimental.
	MultipleValuesSetting interface{} `field:"optional" json:"multipleValuesSetting" yaml:"multipleValuesSetting"`
	// obfuscation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#obfuscation_setting TfSlot#obfuscation_setting}
	// Experimental.
	ObfuscationSetting interface{} `field:"optional" json:"obfuscationSetting" yaml:"obfuscationSetting"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#region TfSlot#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#slot_type_id TfSlot#slot_type_id}.
	// Experimental.
	SlotTypeId *string `field:"optional" json:"slotTypeId" yaml:"slotTypeId"`
	// sub_slot_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#sub_slot_setting TfSlot#sub_slot_setting}
	// Experimental.
	SubSlotSetting interface{} `field:"optional" json:"subSlotSetting" yaml:"subSlotSetting"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#timeouts TfSlot#timeouts}
	// Experimental.
	Timeouts *TfSlot_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// value_elicitation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#value_elicitation_setting TfSlot#value_elicitation_setting}
	// Experimental.
	ValueElicitationSetting interface{} `field:"optional" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
}

