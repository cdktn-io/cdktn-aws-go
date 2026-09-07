package bedrockagentcore


// Experimental.
type AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#type AwsGatewayTarget#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#description AwsGatewayTarget#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#items AwsGatewayTarget#items}
	// Experimental.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#property AwsGatewayTarget#property}
	// Experimental.
	Property interface{} `field:"optional" json:"property" yaml:"property"`
}

