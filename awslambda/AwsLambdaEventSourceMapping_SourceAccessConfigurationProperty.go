package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_SourceAccessConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#type AwsLambdaEventSourceMapping#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#uri AwsLambdaEventSourceMapping#uri}.
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

