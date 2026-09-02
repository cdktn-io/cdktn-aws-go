package awsbedrockagents


// Experimental.
type TfDataSource_SemanticChunkingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#breakpoint_percentile_threshold TfDataSource#breakpoint_percentile_threshold}.
	// Experimental.
	BreakpointPercentileThreshold *float64 `field:"required" json:"breakpointPercentileThreshold" yaml:"breakpointPercentileThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#buffer_size TfDataSource#buffer_size}.
	// Experimental.
	BufferSize *float64 `field:"required" json:"bufferSize" yaml:"bufferSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#max_token TfDataSource#max_token}.
	// Experimental.
	MaxToken *float64 `field:"required" json:"maxToken" yaml:"maxToken"`
}

