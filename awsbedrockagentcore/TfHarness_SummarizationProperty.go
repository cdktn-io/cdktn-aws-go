package awsbedrockagentcore


// Experimental.
type TfHarness_SummarizationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#preserve_recent_messages TfHarness#preserve_recent_messages}.
	// Experimental.
	PreserveRecentMessages *float64 `field:"optional" json:"preserveRecentMessages" yaml:"preserveRecentMessages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#summarization_system_prompt TfHarness#summarization_system_prompt}.
	// Experimental.
	SummarizationSystemPrompt *string `field:"optional" json:"summarizationSystemPrompt" yaml:"summarizationSystemPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#summary_ratio TfHarness#summary_ratio}.
	// Experimental.
	SummaryRatio *float64 `field:"optional" json:"summaryRatio" yaml:"summaryRatio"`
}

