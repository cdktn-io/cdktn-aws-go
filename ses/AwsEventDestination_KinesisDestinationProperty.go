package ses


// Experimental.
type AwsEventDestination_KinesisDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#role_arn AwsEventDestination#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#stream_arn AwsEventDestination#stream_arn}.
	// Experimental.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

