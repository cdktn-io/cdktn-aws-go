package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_InputSchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#type AwsBedrockagentcoreGatewayTarget#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#description AwsBedrockagentcoreGatewayTarget#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#items AwsBedrockagentcoreGatewayTarget#items}
	// Experimental.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#property AwsBedrockagentcoreGatewayTarget#property}
	// Experimental.
	Property interface{} `field:"optional" json:"property" yaml:"property"`
}

