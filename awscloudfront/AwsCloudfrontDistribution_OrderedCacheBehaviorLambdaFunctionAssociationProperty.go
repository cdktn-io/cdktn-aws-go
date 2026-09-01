package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_OrderedCacheBehaviorLambdaFunctionAssociationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#event_type AwsCloudfrontDistribution#event_type}.
	// Experimental.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#lambda_arn AwsCloudfrontDistribution#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#include_body AwsCloudfrontDistribution#include_body}.
	// Experimental.
	IncludeBody interface{} `field:"optional" json:"includeBody" yaml:"includeBody"`
}

