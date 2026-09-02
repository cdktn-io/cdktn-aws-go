package awsbedrockagentcore


// Experimental.
type TfAgentRuntime_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#bucket TfAgentRuntime#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#prefix TfAgentRuntime#prefix}.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#version_id TfAgentRuntime#version_id}.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

