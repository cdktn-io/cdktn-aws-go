package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_CredentialProviderConfigurationProperty struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#api_key AwsBedrockagentcoreGatewayTarget#api_key}
	// Experimental.
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// caller_iam_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#caller_iam_credentials AwsBedrockagentcoreGatewayTarget#caller_iam_credentials}
	// Experimental.
	CallerIamCredentials interface{} `field:"optional" json:"callerIamCredentials" yaml:"callerIamCredentials"`
	// gateway_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#gateway_iam_role AwsBedrockagentcoreGatewayTarget#gateway_iam_role}
	// Experimental.
	GatewayIamRole interface{} `field:"optional" json:"gatewayIamRole" yaml:"gatewayIamRole"`
	// jwt_passthrough block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#jwt_passthrough AwsBedrockagentcoreGatewayTarget#jwt_passthrough}
	// Experimental.
	JwtPassthrough interface{} `field:"optional" json:"jwtPassthrough" yaml:"jwtPassthrough"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#oauth AwsBedrockagentcoreGatewayTarget#oauth}
	// Experimental.
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

