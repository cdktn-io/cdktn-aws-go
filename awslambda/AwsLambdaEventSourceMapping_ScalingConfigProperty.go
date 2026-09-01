package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_ScalingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_concurrency AwsLambdaEventSourceMapping#maximum_concurrency}.
	// Experimental.
	MaximumConcurrency *float64 `field:"optional" json:"maximumConcurrency" yaml:"maximumConcurrency"`
}

