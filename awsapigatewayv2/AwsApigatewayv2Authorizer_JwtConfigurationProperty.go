package awsapigatewayv2


// Experimental.
type AwsApigatewayv2Authorizer_JwtConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#audience AwsApigatewayv2Authorizer#audience}.
	// Experimental.
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#issuer AwsApigatewayv2Authorizer#issuer}.
	// Experimental.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

