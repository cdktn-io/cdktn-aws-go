package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_ToolOverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#method TfGatewayTarget#method}.
	// Experimental.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#name TfGatewayTarget#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#path TfGatewayTarget#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#description TfGatewayTarget#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

