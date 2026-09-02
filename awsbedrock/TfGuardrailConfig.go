package awsbedrock

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGuardrailConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#blocked_input_messaging TfGuardrail#blocked_input_messaging}.
	// Experimental.
	BlockedInputMessaging *string `field:"required" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#blocked_outputs_messaging TfGuardrail#blocked_outputs_messaging}.
	// Experimental.
	BlockedOutputsMessaging *string `field:"required" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#name TfGuardrail#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// content_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#content_policy_config TfGuardrail#content_policy_config}
	// Experimental.
	ContentPolicyConfig interface{} `field:"optional" json:"contentPolicyConfig" yaml:"contentPolicyConfig"`
	// contextual_grounding_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#contextual_grounding_policy_config TfGuardrail#contextual_grounding_policy_config}
	// Experimental.
	ContextualGroundingPolicyConfig interface{} `field:"optional" json:"contextualGroundingPolicyConfig" yaml:"contextualGroundingPolicyConfig"`
	// cross_region_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#cross_region_config TfGuardrail#cross_region_config}
	// Experimental.
	CrossRegionConfig interface{} `field:"optional" json:"crossRegionConfig" yaml:"crossRegionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#description TfGuardrail#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#kms_key_arn TfGuardrail#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#region TfGuardrail#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sensitive_information_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#sensitive_information_policy_config TfGuardrail#sensitive_information_policy_config}
	// Experimental.
	SensitiveInformationPolicyConfig interface{} `field:"optional" json:"sensitiveInformationPolicyConfig" yaml:"sensitiveInformationPolicyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#tags TfGuardrail#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#timeouts TfGuardrail#timeouts}
	// Experimental.
	Timeouts *TfGuardrail_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// topic_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#topic_policy_config TfGuardrail#topic_policy_config}
	// Experimental.
	TopicPolicyConfig interface{} `field:"optional" json:"topicPolicyConfig" yaml:"topicPolicyConfig"`
	// word_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#word_policy_config TfGuardrail#word_policy_config}
	// Experimental.
	WordPolicyConfig interface{} `field:"optional" json:"wordPolicyConfig" yaml:"wordPolicyConfig"`
}

