package awsapigatewayv2


// Experimental.
type AwsApigatewayv2DomainName_MutualTlsAuthenticationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_domain_name#truststore_uri AwsApigatewayv2DomainName#truststore_uri}.
	// Experimental.
	TruststoreUri *string `field:"required" json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_domain_name#truststore_version AwsApigatewayv2DomainName#truststore_version}.
	// Experimental.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

