package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_McpServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#endpoint TfGatewayTarget#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#listing_mode TfGatewayTarget#listing_mode}.
	// Experimental.
	ListingMode *string `field:"optional" json:"listingMode" yaml:"listingMode"`
	// mcp_tool_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#mcp_tool_schema TfGatewayTarget#mcp_tool_schema}
	// Experimental.
	McpToolSchema interface{} `field:"optional" json:"mcpToolSchema" yaml:"mcpToolSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#resource_priority TfGatewayTarget#resource_priority}.
	// Experimental.
	ResourcePriority *float64 `field:"optional" json:"resourcePriority" yaml:"resourcePriority"`
}

