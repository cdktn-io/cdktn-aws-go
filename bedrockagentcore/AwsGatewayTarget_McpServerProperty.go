package bedrockagentcore


// Experimental.
type AwsGatewayTarget_McpServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#endpoint AwsGatewayTarget#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#listing_mode AwsGatewayTarget#listing_mode}.
	// Experimental.
	ListingMode *string `field:"optional" json:"listingMode" yaml:"listingMode"`
	// mcp_tool_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#mcp_tool_schema AwsGatewayTarget#mcp_tool_schema}
	// Experimental.
	McpToolSchema interface{} `field:"optional" json:"mcpToolSchema" yaml:"mcpToolSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#resource_priority AwsGatewayTarget#resource_priority}.
	// Experimental.
	ResourcePriority *float64 `field:"optional" json:"resourcePriority" yaml:"resourcePriority"`
}

