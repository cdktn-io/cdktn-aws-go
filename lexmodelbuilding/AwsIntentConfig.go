package lexmodelbuilding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntentConfig struct {
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
	// fulfillment_activity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#fulfillment_activity AwsIntent#fulfillment_activity}
	// Experimental.
	FulfillmentActivity *AwsIntent_FulfillmentActivityProperty `field:"required" json:"fulfillmentActivity" yaml:"fulfillmentActivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#name AwsIntent#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// conclusion_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#conclusion_statement AwsIntent#conclusion_statement}
	// Experimental.
	ConclusionStatement *AwsIntent_ConclusionStatementProperty `field:"optional" json:"conclusionStatement" yaml:"conclusionStatement"`
	// confirmation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#confirmation_prompt AwsIntent#confirmation_prompt}
	// Experimental.
	ConfirmationPrompt *AwsIntent_ConfirmationPromptProperty `field:"optional" json:"confirmationPrompt" yaml:"confirmationPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#create_version AwsIntent#create_version}.
	// Experimental.
	CreateVersion interface{} `field:"optional" json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#description AwsIntent#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dialog_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#dialog_code_hook AwsIntent#dialog_code_hook}
	// Experimental.
	DialogCodeHook *AwsIntent_DialogCodeHookProperty `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// follow_up_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#follow_up_prompt AwsIntent#follow_up_prompt}
	// Experimental.
	FollowUpPrompt *AwsIntent_FollowUpPromptProperty `field:"optional" json:"followUpPrompt" yaml:"followUpPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#id AwsIntent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#parent_intent_signature AwsIntent#parent_intent_signature}.
	// Experimental.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#region AwsIntent#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#rejection_statement AwsIntent#rejection_statement}
	// Experimental.
	RejectionStatement *AwsIntent_RejectionStatementProperty `field:"optional" json:"rejectionStatement" yaml:"rejectionStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#sample_utterances AwsIntent#sample_utterances}.
	// Experimental.
	SampleUtterances *[]*string `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// slot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot AwsIntent#slot}
	// Experimental.
	Slot interface{} `field:"optional" json:"slot" yaml:"slot"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#timeouts AwsIntent#timeouts}
	// Experimental.
	Timeouts *AwsIntent_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

