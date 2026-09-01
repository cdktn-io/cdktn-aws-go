package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#name AwsBedrockagentcoreGatewayTarget#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#type AwsBedrockagentcoreGatewayTarget#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#description AwsBedrockagentcoreGatewayTarget#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#items_json AwsBedrockagentcoreGatewayTarget#items_json}.
	// Experimental.
	ItemsJson *string `field:"optional" json:"itemsJson" yaml:"itemsJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#properties_json AwsBedrockagentcoreGatewayTarget#properties_json}.
	// Experimental.
	PropertiesJson *string `field:"optional" json:"propertiesJson" yaml:"propertiesJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#required AwsBedrockagentcoreGatewayTarget#required}.
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

