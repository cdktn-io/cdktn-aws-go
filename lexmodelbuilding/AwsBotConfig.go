package lexmodelbuilding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBotConfig struct {
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
	// abort_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#abort_statement AwsBot#abort_statement}
	// Experimental.
	AbortStatement *AwsBot_AbortStatementProperty `field:"required" json:"abortStatement" yaml:"abortStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#child_directed AwsBot#child_directed}.
	// Experimental.
	ChildDirected interface{} `field:"required" json:"childDirected" yaml:"childDirected"`
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#intent AwsBot#intent}
	// Experimental.
	Intent interface{} `field:"required" json:"intent" yaml:"intent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#name AwsBot#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// clarification_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#clarification_prompt AwsBot#clarification_prompt}
	// Experimental.
	ClarificationPrompt *AwsBot_ClarificationPromptProperty `field:"optional" json:"clarificationPrompt" yaml:"clarificationPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#create_version AwsBot#create_version}.
	// Experimental.
	CreateVersion interface{} `field:"optional" json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#description AwsBot#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#detect_sentiment AwsBot#detect_sentiment}.
	// Experimental.
	DetectSentiment interface{} `field:"optional" json:"detectSentiment" yaml:"detectSentiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#enable_model_improvements AwsBot#enable_model_improvements}.
	// Experimental.
	EnableModelImprovements interface{} `field:"optional" json:"enableModelImprovements" yaml:"enableModelImprovements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#id AwsBot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#idle_session_ttl_in_seconds AwsBot#idle_session_ttl_in_seconds}.
	// Experimental.
	IdleSessionTtlInSeconds *float64 `field:"optional" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#locale AwsBot#locale}.
	// Experimental.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#nlu_intent_confidence_threshold AwsBot#nlu_intent_confidence_threshold}.
	// Experimental.
	NluIntentConfidenceThreshold *float64 `field:"optional" json:"nluIntentConfidenceThreshold" yaml:"nluIntentConfidenceThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#process_behavior AwsBot#process_behavior}.
	// Experimental.
	ProcessBehavior *string `field:"optional" json:"processBehavior" yaml:"processBehavior"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#region AwsBot#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#timeouts AwsBot#timeouts}
	// Experimental.
	Timeouts *AwsBot_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#voice_id AwsBot#voice_id}.
	// Experimental.
	VoiceId *string `field:"optional" json:"voiceId" yaml:"voiceId"`
}

