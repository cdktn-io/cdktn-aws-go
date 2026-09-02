package awscloudfront


// Experimental.
type TfMultitenantDistribution_CacheBehaviorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#path_pattern TfMultitenantDistribution#path_pattern}.
	// Experimental.
	PathPattern *string `field:"required" json:"pathPattern" yaml:"pathPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#target_origin_id TfMultitenantDistribution#target_origin_id}.
	// Experimental.
	TargetOriginId *string `field:"required" json:"targetOriginId" yaml:"targetOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#viewer_protocol_policy TfMultitenantDistribution#viewer_protocol_policy}.
	// Experimental.
	ViewerProtocolPolicy *string `field:"required" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// allowed_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#allowed_methods TfMultitenantDistribution#allowed_methods}
	// Experimental.
	AllowedMethods interface{} `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cache_policy_id TfMultitenantDistribution#cache_policy_id}.
	// Experimental.
	CachePolicyId *string `field:"optional" json:"cachePolicyId" yaml:"cachePolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#compress TfMultitenantDistribution#compress}.
	// Experimental.
	Compress interface{} `field:"optional" json:"compress" yaml:"compress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#field_level_encryption_id TfMultitenantDistribution#field_level_encryption_id}.
	// Experimental.
	FieldLevelEncryptionId *string `field:"optional" json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#function_association TfMultitenantDistribution#function_association}
	// Experimental.
	FunctionAssociation interface{} `field:"optional" json:"functionAssociation" yaml:"functionAssociation"`
	// lambda_function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#lambda_function_association TfMultitenantDistribution#lambda_function_association}
	// Experimental.
	LambdaFunctionAssociation interface{} `field:"optional" json:"lambdaFunctionAssociation" yaml:"lambdaFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_request_policy_id TfMultitenantDistribution#origin_request_policy_id}.
	// Experimental.
	OriginRequestPolicyId *string `field:"optional" json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#realtime_log_config_arn TfMultitenantDistribution#realtime_log_config_arn}.
	// Experimental.
	RealtimeLogConfigArn *string `field:"optional" json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#response_headers_policy_id TfMultitenantDistribution#response_headers_policy_id}.
	// Experimental.
	ResponseHeadersPolicyId *string `field:"optional" json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// trusted_key_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#trusted_key_groups TfMultitenantDistribution#trusted_key_groups}
	// Experimental.
	TrustedKeyGroups interface{} `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
}

