package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_OauthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#provider_arn AwsBedrockagentcoreHarness#provider_arn}.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#scopes AwsBedrockagentcoreHarness#scopes}.
	// Experimental.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#custom_parameters AwsBedrockagentcoreHarness#custom_parameters}.
	// Experimental.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#default_return_url AwsBedrockagentcoreHarness#default_return_url}.
	// Experimental.
	DefaultReturnUrl *string `field:"optional" json:"defaultReturnUrl" yaml:"defaultReturnUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#grant_type AwsBedrockagentcoreHarness#grant_type}.
	// Experimental.
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
}

