package awsbedrockagentcore


// Experimental.
type TfHarness_TruncationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#sliding_window TfHarness#sliding_window}.
	// Experimental.
	SlidingWindow interface{} `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#summarization TfHarness#summarization}.
	// Experimental.
	Summarization interface{} `field:"optional" json:"summarization" yaml:"summarization"`
}

