package awsappflow


// Experimental.
type TfFlow_SnowflakeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#intermediate_bucket_name TfFlow#intermediate_bucket_name}.
	// Experimental.
	IntermediateBucketName *string `field:"required" json:"intermediateBucketName" yaml:"intermediateBucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object TfFlow#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_prefix TfFlow#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#error_handling_config TfFlow#error_handling_config}
	// Experimental.
	ErrorHandlingConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfigProperty `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

