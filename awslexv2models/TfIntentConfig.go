package awslexv2models

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#bot_id TfIntent#bot_id}.
	// Experimental.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#bot_version TfIntent#bot_version}.
	// Experimental.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#locale_id TfIntent#locale_id}.
	// Experimental.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#name TfIntent#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// closing_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#closing_setting TfIntent#closing_setting}
	// Experimental.
	ClosingSetting interface{} `field:"optional" json:"closingSetting" yaml:"closingSetting"`
	// confirmation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#confirmation_setting TfIntent#confirmation_setting}
	// Experimental.
	ConfirmationSetting interface{} `field:"optional" json:"confirmationSetting" yaml:"confirmationSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#description TfIntent#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dialog_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#dialog_code_hook TfIntent#dialog_code_hook}
	// Experimental.
	DialogCodeHook interface{} `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// fulfillment_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#fulfillment_code_hook TfIntent#fulfillment_code_hook}
	// Experimental.
	FulfillmentCodeHook interface{} `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// initial_response_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#initial_response_setting TfIntent#initial_response_setting}
	// Experimental.
	InitialResponseSetting interface{} `field:"optional" json:"initialResponseSetting" yaml:"initialResponseSetting"`
	// input_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#input_context TfIntent#input_context}
	// Experimental.
	InputContext interface{} `field:"optional" json:"inputContext" yaml:"inputContext"`
	// kendra_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#kendra_configuration TfIntent#kendra_configuration}
	// Experimental.
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// output_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#output_context TfIntent#output_context}
	// Experimental.
	OutputContext interface{} `field:"optional" json:"outputContext" yaml:"outputContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#parent_intent_signature TfIntent#parent_intent_signature}.
	// Experimental.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// qna_intent_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#qna_intent_configuration TfIntent#qna_intent_configuration}
	// Experimental.
	QnaIntentConfiguration interface{} `field:"optional" json:"qnaIntentConfiguration" yaml:"qnaIntentConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#region TfIntent#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sample_utterance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#sample_utterance TfIntent#sample_utterance}
	// Experimental.
	SampleUtterance interface{} `field:"optional" json:"sampleUtterance" yaml:"sampleUtterance"`
	// slot_priority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#slot_priority TfIntent#slot_priority}
	// Experimental.
	SlotPriority interface{} `field:"optional" json:"slotPriority" yaml:"slotPriority"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#timeouts TfIntent#timeouts}
	// Experimental.
	Timeouts *TfIntent_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

