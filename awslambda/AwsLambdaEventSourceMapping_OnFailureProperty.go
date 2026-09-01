package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_OnFailureProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#destination_arn AwsLambdaEventSourceMapping#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
}

