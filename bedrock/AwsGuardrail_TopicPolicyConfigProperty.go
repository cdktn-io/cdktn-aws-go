package bedrock


// Experimental.
type AwsGuardrail_TopicPolicyConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#tier_config AwsGuardrail#tier_config}.
	// Experimental.
	TierConfig interface{} `field:"optional" json:"tierConfig" yaml:"tierConfig"`
	// topics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#topics_config AwsGuardrail#topics_config}
	// Experimental.
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
}

