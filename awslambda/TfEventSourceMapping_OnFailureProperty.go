package awslambda


// Experimental.
type TfEventSourceMapping_OnFailureProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#destination_arn TfEventSourceMapping#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
}

