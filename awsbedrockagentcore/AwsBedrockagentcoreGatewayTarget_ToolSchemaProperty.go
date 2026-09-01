package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_ToolSchemaProperty struct {
	// inline_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#inline_payload AwsBedrockagentcoreGatewayTarget#inline_payload}
	// Experimental.
	InlinePayload interface{} `field:"optional" json:"inlinePayload" yaml:"inlinePayload"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#s3 AwsBedrockagentcoreGatewayTarget#s3}
	// Experimental.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

