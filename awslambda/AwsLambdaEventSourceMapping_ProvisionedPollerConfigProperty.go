package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_pollers AwsLambdaEventSourceMapping#maximum_pollers}.
	// Experimental.
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#minimum_pollers AwsLambdaEventSourceMapping#minimum_pollers}.
	// Experimental.
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#poller_group_name AwsLambdaEventSourceMapping#poller_group_name}.
	// Experimental.
	PollerGroupName *string `field:"optional" json:"pollerGroupName" yaml:"pollerGroupName"`
}

