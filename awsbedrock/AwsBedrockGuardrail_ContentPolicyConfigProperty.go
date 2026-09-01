package awsbedrock


// Experimental.
type AwsBedrockGuardrail_ContentPolicyConfigProperty struct {
	// filters_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#filters_config AwsBedrockGuardrail#filters_config}
	// Experimental.
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#tier_config AwsBedrockGuardrail#tier_config}.
	// Experimental.
	TierConfig interface{} `field:"optional" json:"tierConfig" yaml:"tierConfig"`
}

