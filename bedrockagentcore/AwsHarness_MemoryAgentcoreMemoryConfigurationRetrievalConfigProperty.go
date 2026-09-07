package bedrockagentcore


// Experimental.
type AwsHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#map_block_key AwsHarness#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#relevance_score AwsHarness#relevance_score}.
	// Experimental.
	RelevanceScore *float64 `field:"optional" json:"relevanceScore" yaml:"relevanceScore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#strategy_id AwsHarness#strategy_id}.
	// Experimental.
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#top_k AwsHarness#top_k}.
	// Experimental.
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
}

