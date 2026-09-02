package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_ToolFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#filter_path TfGatewayTarget#filter_path}.
	// Experimental.
	FilterPath *string `field:"required" json:"filterPath" yaml:"filterPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#methods TfGatewayTarget#methods}.
	// Experimental.
	Methods *[]*string `field:"required" json:"methods" yaml:"methods"`
}

