package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_DefaultCacheBehaviorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#allowed_methods AwsCloudfrontDistribution#allowed_methods}.
	// Experimental.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cached_methods AwsCloudfrontDistribution#cached_methods}.
	// Experimental.
	CachedMethods *[]*string `field:"required" json:"cachedMethods" yaml:"cachedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#target_origin_id AwsCloudfrontDistribution#target_origin_id}.
	// Experimental.
	TargetOriginId *string `field:"required" json:"targetOriginId" yaml:"targetOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#viewer_protocol_policy AwsCloudfrontDistribution#viewer_protocol_policy}.
	// Experimental.
	ViewerProtocolPolicy *string `field:"required" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cache_policy_id AwsCloudfrontDistribution#cache_policy_id}.
	// Experimental.
	CachePolicyId *string `field:"optional" json:"cachePolicyId" yaml:"cachePolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#compress AwsCloudfrontDistribution#compress}.
	// Experimental.
	Compress interface{} `field:"optional" json:"compress" yaml:"compress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#default_ttl AwsCloudfrontDistribution#default_ttl}.
	// Experimental.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#field_level_encryption_id AwsCloudfrontDistribution#field_level_encryption_id}.
	// Experimental.
	FieldLevelEncryptionId *string `field:"optional" json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// forwarded_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#forwarded_values AwsCloudfrontDistribution#forwarded_values}
	// Experimental.
	ForwardedValues *AwsCloudfrontDistribution_DefaultCacheBehaviorForwardedValuesProperty `field:"optional" json:"forwardedValues" yaml:"forwardedValues"`
	// function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#function_association AwsCloudfrontDistribution#function_association}
	// Experimental.
	FunctionAssociation interface{} `field:"optional" json:"functionAssociation" yaml:"functionAssociation"`
	// grpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#grpc_config AwsCloudfrontDistribution#grpc_config}
	// Experimental.
	GrpcConfig *AwsCloudfrontDistribution_DefaultCacheBehaviorGrpcConfigProperty `field:"optional" json:"grpcConfig" yaml:"grpcConfig"`
	// lambda_function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#lambda_function_association AwsCloudfrontDistribution#lambda_function_association}
	// Experimental.
	LambdaFunctionAssociation interface{} `field:"optional" json:"lambdaFunctionAssociation" yaml:"lambdaFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#max_ttl AwsCloudfrontDistribution#max_ttl}.
	// Experimental.
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#min_ttl AwsCloudfrontDistribution#min_ttl}.
	// Experimental.
	MinTtl *float64 `field:"optional" json:"minTtl" yaml:"minTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_request_policy_id AwsCloudfrontDistribution#origin_request_policy_id}.
	// Experimental.
	OriginRequestPolicyId *string `field:"optional" json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#realtime_log_config_arn AwsCloudfrontDistribution#realtime_log_config_arn}.
	// Experimental.
	RealtimeLogConfigArn *string `field:"optional" json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#response_headers_policy_id AwsCloudfrontDistribution#response_headers_policy_id}.
	// Experimental.
	ResponseHeadersPolicyId *string `field:"optional" json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#smooth_streaming AwsCloudfrontDistribution#smooth_streaming}.
	// Experimental.
	SmoothStreaming interface{} `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#trusted_key_groups AwsCloudfrontDistribution#trusted_key_groups}.
	// Experimental.
	TrustedKeyGroups *[]*string `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#trusted_signers AwsCloudfrontDistribution#trusted_signers}.
	// Experimental.
	TrustedSigners *[]*string `field:"optional" json:"trustedSigners" yaml:"trustedSigners"`
}

