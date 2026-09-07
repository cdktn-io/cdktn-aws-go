package bedrockagentcore


// Experimental.
type AwsGatewayRule_IamPrincipalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#arn AwsGatewayRule#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#operator AwsGatewayRule#operator}.
	// Experimental.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

