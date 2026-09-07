package opensearchserverless


// Experimental.
type AwsSecurityConfig_IamFederationOptionsProperty struct {
	// Group attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#group_attribute AwsSecurityConfig#group_attribute}
	// Experimental.
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// User attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#user_attribute AwsSecurityConfig#user_attribute}
	// Experimental.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

