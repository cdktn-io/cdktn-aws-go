package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#event_type AwsCloudfrontMultitenantDistribution#event_type}.
	// Experimental.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#lambda_function_arn AwsCloudfrontMultitenantDistribution#lambda_function_arn}.
	// Experimental.
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#include_body AwsCloudfrontMultitenantDistribution#include_body}.
	// Experimental.
	IncludeBody interface{} `field:"optional" json:"includeBody" yaml:"includeBody"`
}

