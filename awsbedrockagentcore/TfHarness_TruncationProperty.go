package awsbedrockagentcore


// Experimental.
type TfHarness_TruncationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#config TfHarness#config}.
	// Experimental.
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#strategy TfHarness#strategy}.
	// Experimental.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

