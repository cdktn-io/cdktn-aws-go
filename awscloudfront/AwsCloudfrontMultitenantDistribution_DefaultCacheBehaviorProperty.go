package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#target_origin_id AwsCloudfrontMultitenantDistribution#target_origin_id}.
	// Experimental.
	TargetOriginId *string `field:"required" json:"targetOriginId" yaml:"targetOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#viewer_protocol_policy AwsCloudfrontMultitenantDistribution#viewer_protocol_policy}.
	// Experimental.
	ViewerProtocolPolicy *string `field:"required" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// allowed_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#allowed_methods AwsCloudfrontMultitenantDistribution#allowed_methods}
	// Experimental.
	AllowedMethods interface{} `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cache_policy_id AwsCloudfrontMultitenantDistribution#cache_policy_id}.
	// Experimental.
	CachePolicyId *string `field:"optional" json:"cachePolicyId" yaml:"cachePolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#compress AwsCloudfrontMultitenantDistribution#compress}.
	// Experimental.
	Compress interface{} `field:"optional" json:"compress" yaml:"compress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#field_level_encryption_id AwsCloudfrontMultitenantDistribution#field_level_encryption_id}.
	// Experimental.
	FieldLevelEncryptionId *string `field:"optional" json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#function_association AwsCloudfrontMultitenantDistribution#function_association}
	// Experimental.
	FunctionAssociation interface{} `field:"optional" json:"functionAssociation" yaml:"functionAssociation"`
	// lambda_function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#lambda_function_association AwsCloudfrontMultitenantDistribution#lambda_function_association}
	// Experimental.
	LambdaFunctionAssociation interface{} `field:"optional" json:"lambdaFunctionAssociation" yaml:"lambdaFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_request_policy_id AwsCloudfrontMultitenantDistribution#origin_request_policy_id}.
	// Experimental.
	OriginRequestPolicyId *string `field:"optional" json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#realtime_log_config_arn AwsCloudfrontMultitenantDistribution#realtime_log_config_arn}.
	// Experimental.
	RealtimeLogConfigArn *string `field:"optional" json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#response_headers_policy_id AwsCloudfrontMultitenantDistribution#response_headers_policy_id}.
	// Experimental.
	ResponseHeadersPolicyId *string `field:"optional" json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// trusted_key_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#trusted_key_groups AwsCloudfrontMultitenantDistribution#trusted_key_groups}
	// Experimental.
	TrustedKeyGroups interface{} `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
}

