package awsopensearchserverless


// Experimental.
type AwsOpensearchserverlessSecurityConfig_SamlOptionsProperty struct {
	// The XML IdP metadata file generated from your identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#metadata AwsOpensearchserverlessSecurityConfig#metadata}
	// Experimental.
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// Group attribute for this SAML integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#group_attribute AwsOpensearchserverlessSecurityConfig#group_attribute}
	// Experimental.
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#session_timeout AwsOpensearchserverlessSecurityConfig#session_timeout}
	// Experimental.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// User attribute for this SAML integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#user_attribute AwsOpensearchserverlessSecurityConfig#user_attribute}
	// Experimental.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

