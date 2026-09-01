package awsopensearchserverless


// Experimental.
type AwsOpensearchserverlessSecurityConfig_IamIdentityCenterOptionsProperty struct {
	// Instance ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#instance_arn AwsOpensearchserverlessSecurityConfig#instance_arn}
	// Experimental.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Group attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#group_attribute AwsOpensearchserverlessSecurityConfig#group_attribute}
	// Experimental.
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// User attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#user_attribute AwsOpensearchserverlessSecurityConfig#user_attribute}
	// Experimental.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

