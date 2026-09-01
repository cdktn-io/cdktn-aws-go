package awsapigateway


// Experimental.
type AwsApiGatewayDomainName_MutualTlsAuthenticationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#truststore_uri AwsApiGatewayDomainName#truststore_uri}.
	// Experimental.
	TruststoreUri *string `field:"required" json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#truststore_version AwsApiGatewayDomainName#truststore_version}.
	// Experimental.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

