package awsbedrock


// Experimental.
type TfGuardrail_TopicPolicyConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#tier_config TfGuardrail#tier_config}.
	// Experimental.
	TierConfig interface{} `field:"optional" json:"tierConfig" yaml:"tierConfig"`
	// topics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#topics_config TfGuardrail#topics_config}
	// Experimental.
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
}

