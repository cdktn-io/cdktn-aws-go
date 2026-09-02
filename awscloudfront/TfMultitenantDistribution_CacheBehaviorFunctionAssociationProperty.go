package awscloudfront


// Experimental.
type TfMultitenantDistribution_CacheBehaviorFunctionAssociationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#event_type TfMultitenantDistribution#event_type}.
	// Experimental.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#function_arn TfMultitenantDistribution#function_arn}.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

