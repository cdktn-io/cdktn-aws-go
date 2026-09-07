package appflow


// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_name AwsFlow#bucket_name}.
	// Experimental.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_prefix AwsFlow#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#fail_on_first_destination_error AwsFlow#fail_on_first_destination_error}.
	// Experimental.
	FailOnFirstDestinationError interface{} `field:"optional" json:"failOnFirstDestinationError" yaml:"failOnFirstDestinationError"`
}

