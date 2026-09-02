package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_TargetConfigurationProperty struct {
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#http TfGatewayTarget#http}
	// Experimental.
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// mcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#mcp TfGatewayTarget#mcp}
	// Experimental.
	Mcp interface{} `field:"optional" json:"mcp" yaml:"mcp"`
}

