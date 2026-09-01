package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#consumer_group_id AwsLambdaEventSourceMapping#consumer_group_id}.
	// Experimental.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// schema_registry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#schema_registry_config AwsLambdaEventSourceMapping#schema_registry_config}
	// Experimental.
	SchemaRegistryConfig *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty `field:"optional" json:"schemaRegistryConfig" yaml:"schemaRegistryConfig"`
}

