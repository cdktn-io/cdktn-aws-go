package bedrockagentcore


// Experimental.
type AwsGatewayTarget_ApiKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#provider_arn AwsGatewayTarget#provider_arn}.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#credential_location AwsGatewayTarget#credential_location}.
	// Experimental.
	CredentialLocation *string `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#credential_parameter_name AwsGatewayTarget#credential_parameter_name}.
	// Experimental.
	CredentialParameterName *string `field:"optional" json:"credentialParameterName" yaml:"credentialParameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#credential_prefix AwsGatewayTarget#credential_prefix}.
	// Experimental.
	CredentialPrefix *string `field:"optional" json:"credentialPrefix" yaml:"credentialPrefix"`
}

