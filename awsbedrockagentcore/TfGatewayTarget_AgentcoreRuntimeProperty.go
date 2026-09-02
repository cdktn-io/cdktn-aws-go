package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_AgentcoreRuntimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#arn TfGatewayTarget#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#qualifier TfGatewayTarget#qualifier}.
	// Experimental.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

