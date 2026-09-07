package bedrock


// Experimental.
type AwsModelInvocationLoggingConfiguration_CloudwatchConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#log_group_name AwsModelInvocationLoggingConfiguration#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#role_arn AwsModelInvocationLoggingConfiguration#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// large_data_delivery_s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#large_data_delivery_s3_config AwsModelInvocationLoggingConfiguration#large_data_delivery_s3_config}
	// Experimental.
	LargeDataDeliveryS3Config interface{} `field:"optional" json:"largeDataDeliveryS3Config" yaml:"largeDataDeliveryS3Config"`
}

