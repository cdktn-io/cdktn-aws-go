package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_RemoteMcpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#url AwsBedrockagentcoreHarness#url}.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#headers AwsBedrockagentcoreHarness#headers}.
	// Experimental.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
}

