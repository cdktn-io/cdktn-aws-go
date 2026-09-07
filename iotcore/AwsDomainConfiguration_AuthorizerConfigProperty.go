package iotcore


// Experimental.
type AwsDomainConfiguration_AuthorizerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#allow_authorizer_override AwsDomainConfiguration#allow_authorizer_override}.
	// Experimental.
	AllowAuthorizerOverride interface{} `field:"optional" json:"allowAuthorizerOverride" yaml:"allowAuthorizerOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#default_authorizer_name AwsDomainConfiguration#default_authorizer_name}.
	// Experimental.
	DefaultAuthorizerName *string `field:"optional" json:"defaultAuthorizerName" yaml:"defaultAuthorizerName"`
}

