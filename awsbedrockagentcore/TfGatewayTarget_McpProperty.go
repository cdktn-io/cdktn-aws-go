package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_McpProperty struct {
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#api_gateway TfGatewayTarget#api_gateway}
	// Experimental.
	ApiGateway interface{} `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#lambda TfGatewayTarget#lambda}
	// Experimental.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// mcp_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#mcp_server TfGatewayTarget#mcp_server}
	// Experimental.
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// open_api_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#open_api_schema TfGatewayTarget#open_api_schema}
	// Experimental.
	OpenApiSchema interface{} `field:"optional" json:"openApiSchema" yaml:"openApiSchema"`
	// smithy_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#smithy_model TfGatewayTarget#smithy_model}
	// Experimental.
	SmithyModel interface{} `field:"optional" json:"smithyModel" yaml:"smithyModel"`
}

