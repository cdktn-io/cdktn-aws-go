package awsapigateway


// Experimental.
type TfDomainName_MutualTlsAuthenticationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#truststore_uri TfDomainName#truststore_uri}.
	// Experimental.
	TruststoreUri *string `field:"required" json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#truststore_version TfDomainName#truststore_version}.
	// Experimental.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

