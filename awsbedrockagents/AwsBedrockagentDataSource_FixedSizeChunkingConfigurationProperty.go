package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_FixedSizeChunkingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#max_tokens AwsBedrockagentDataSource#max_tokens}.
	// Experimental.
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#overlap_percentage AwsBedrockagentDataSource#overlap_percentage}.
	// Experimental.
	OverlapPercentage *float64 `field:"required" json:"overlapPercentage" yaml:"overlapPercentage"`
}

