package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_OauthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#provider_arn AwsBedrockagentcoreGatewayTarget#provider_arn}.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#scopes AwsBedrockagentcoreGatewayTarget#scopes}.
	// Experimental.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#custom_parameters AwsBedrockagentcoreGatewayTarget#custom_parameters}.
	// Experimental.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
	// The URL where the end user's browser is redirected after obtaining the authorization code. Required when grant_type is AUTHORIZATION_CODE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#default_return_url AwsBedrockagentcoreGatewayTarget#default_return_url}
	// Experimental.
	DefaultReturnUrl *string `field:"optional" json:"defaultReturnUrl" yaml:"defaultReturnUrl"`
	// The OAuth grant type. Valid values are AUTHORIZATION_CODE and CLIENT_CREDENTIALS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#grant_type AwsBedrockagentcoreGatewayTarget#grant_type}
	// Experimental.
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
}

