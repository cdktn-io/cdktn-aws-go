package bedrockagentcore


// Experimental.
type AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#description AwsGatewayTarget#description}.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#name AwsGatewayTarget#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#input_schema AwsGatewayTarget#input_schema}
	// Experimental.
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// output_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#output_schema AwsGatewayTarget#output_schema}
	// Experimental.
	OutputSchema interface{} `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

