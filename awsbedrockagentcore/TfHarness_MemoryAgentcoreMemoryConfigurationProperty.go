package awsbedrockagentcore


// Experimental.
type TfHarness_MemoryAgentcoreMemoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#arn TfHarness#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#actor_id TfHarness#actor_id}.
	// Experimental.
	ActorId *string `field:"optional" json:"actorId" yaml:"actorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#messages_count TfHarness#messages_count}.
	// Experimental.
	MessagesCount *float64 `field:"optional" json:"messagesCount" yaml:"messagesCount"`
	// retrieval_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#retrieval_config TfHarness#retrieval_config}
	// Experimental.
	RetrievalConfig interface{} `field:"optional" json:"retrievalConfig" yaml:"retrievalConfig"`
}

