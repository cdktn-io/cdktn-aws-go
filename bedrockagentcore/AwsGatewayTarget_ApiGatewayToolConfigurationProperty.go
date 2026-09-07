package bedrockagentcore


// Experimental.
type AwsGatewayTarget_ApiGatewayToolConfigurationProperty struct {
	// tool_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#tool_filter AwsGatewayTarget#tool_filter}
	// Experimental.
	ToolFilter interface{} `field:"optional" json:"toolFilter" yaml:"toolFilter"`
	// tool_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#tool_override AwsGatewayTarget#tool_override}
	// Experimental.
	ToolOverride interface{} `field:"optional" json:"toolOverride" yaml:"toolOverride"`
}

