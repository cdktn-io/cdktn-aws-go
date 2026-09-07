package apigatewayv2


// Experimental.
type AwsDomainName_MutualTlsAuthenticationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_domain_name#truststore_uri AwsDomainName#truststore_uri}.
	// Experimental.
	TruststoreUri *string `field:"required" json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_domain_name#truststore_version AwsDomainName#truststore_version}.
	// Experimental.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

