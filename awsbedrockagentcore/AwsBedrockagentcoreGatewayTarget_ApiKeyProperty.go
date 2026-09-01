package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_ApiKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#provider_arn AwsBedrockagentcoreGatewayTarget#provider_arn}.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#credential_location AwsBedrockagentcoreGatewayTarget#credential_location}.
	// Experimental.
	CredentialLocation *string `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#credential_parameter_name AwsBedrockagentcoreGatewayTarget#credential_parameter_name}.
	// Experimental.
	CredentialParameterName *string `field:"optional" json:"credentialParameterName" yaml:"credentialParameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#credential_prefix AwsBedrockagentcoreGatewayTarget#credential_prefix}.
	// Experimental.
	CredentialPrefix *string `field:"optional" json:"credentialPrefix" yaml:"credentialPrefix"`
}

