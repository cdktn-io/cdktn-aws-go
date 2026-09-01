package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreMemoryStrategy_TriggerConditionsProperty struct {
	// message_based_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#message_based_trigger AwsBedrockagentcoreMemoryStrategy#message_based_trigger}
	// Experimental.
	MessageBasedTrigger interface{} `field:"optional" json:"messageBasedTrigger" yaml:"messageBasedTrigger"`
	// time_based_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#time_based_trigger AwsBedrockagentcoreMemoryStrategy#time_based_trigger}
	// Experimental.
	TimeBasedTrigger interface{} `field:"optional" json:"timeBasedTrigger" yaml:"timeBasedTrigger"`
	// token_based_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#token_based_trigger AwsBedrockagentcoreMemoryStrategy#token_based_trigger}
	// Experimental.
	TokenBasedTrigger interface{} `field:"optional" json:"tokenBasedTrigger" yaml:"tokenBasedTrigger"`
}

