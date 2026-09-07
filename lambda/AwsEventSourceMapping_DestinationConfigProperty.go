package lambda


// Experimental.
type AwsEventSourceMapping_DestinationConfigProperty struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#on_failure AwsEventSourceMapping#on_failure}
	// Experimental.
	OnFailure *AwsEventSourceMapping_OnFailureProperty `field:"optional" json:"onFailure" yaml:"onFailure"`
}

