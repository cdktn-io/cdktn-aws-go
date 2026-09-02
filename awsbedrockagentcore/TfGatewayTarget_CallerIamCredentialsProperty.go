package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_CallerIamCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#service TfGatewayTarget#service}.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#region TfGatewayTarget#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

