package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_CallerIamCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#service AwsBedrockagentcoreGatewayTarget#service}.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#region AwsBedrockagentcoreGatewayTarget#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

