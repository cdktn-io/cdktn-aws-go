package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_InvocationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#payload_delivery_bucket_name TfMemoryStrategy#payload_delivery_bucket_name}.
	// Experimental.
	PayloadDeliveryBucketName *string `field:"required" json:"payloadDeliveryBucketName" yaml:"payloadDeliveryBucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#topic_arn TfMemoryStrategy#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

