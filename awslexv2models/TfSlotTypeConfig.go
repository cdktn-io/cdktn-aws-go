package awslexv2models

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlotTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#bot_id TfSlotType#bot_id}.
	// Experimental.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#bot_version TfSlotType#bot_version}.
	// Experimental.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#locale_id TfSlotType#locale_id}.
	// Experimental.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#name TfSlotType#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// composite_slot_type_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#composite_slot_type_setting TfSlotType#composite_slot_type_setting}
	// Experimental.
	CompositeSlotTypeSetting interface{} `field:"optional" json:"compositeSlotTypeSetting" yaml:"compositeSlotTypeSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#description TfSlotType#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// external_source_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#external_source_setting TfSlotType#external_source_setting}
	// Experimental.
	ExternalSourceSetting interface{} `field:"optional" json:"externalSourceSetting" yaml:"externalSourceSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#parent_slot_type_signature TfSlotType#parent_slot_type_signature}.
	// Experimental.
	ParentSlotTypeSignature *string `field:"optional" json:"parentSlotTypeSignature" yaml:"parentSlotTypeSignature"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#region TfSlotType#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// slot_type_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#slot_type_values TfSlotType#slot_type_values}
	// Experimental.
	SlotTypeValues interface{} `field:"optional" json:"slotTypeValues" yaml:"slotTypeValues"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#timeouts TfSlotType#timeouts}
	// Experimental.
	Timeouts *TfSlotType_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// value_selection_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#value_selection_setting TfSlotType#value_selection_setting}
	// Experimental.
	ValueSelectionSetting interface{} `field:"optional" json:"valueSelectionSetting" yaml:"valueSelectionSetting"`
}

