package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreMemoryStrategy_SelfManagedConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#historical_context_window_size AwsBedrockagentcoreMemoryStrategy#historical_context_window_size}.
	// Experimental.
	HistoricalContextWindowSize *float64 `field:"optional" json:"historicalContextWindowSize" yaml:"historicalContextWindowSize"`
	// invocation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#invocation_configuration AwsBedrockagentcoreMemoryStrategy#invocation_configuration}
	// Experimental.
	InvocationConfiguration interface{} `field:"optional" json:"invocationConfiguration" yaml:"invocationConfiguration"`
	// trigger_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#trigger_conditions AwsBedrockagentcoreMemoryStrategy#trigger_conditions}
	// Experimental.
	TriggerConditions interface{} `field:"optional" json:"triggerConditions" yaml:"triggerConditions"`
}

