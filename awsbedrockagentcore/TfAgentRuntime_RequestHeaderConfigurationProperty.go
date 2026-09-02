package awsbedrockagentcore


// Experimental.
type TfAgentRuntime_RequestHeaderConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#request_header_allowlist TfAgentRuntime#request_header_allowlist}.
	// Experimental.
	RequestHeaderAllowlist *[]*string `field:"optional" json:"requestHeaderAllowlist" yaml:"requestHeaderAllowlist"`
}

