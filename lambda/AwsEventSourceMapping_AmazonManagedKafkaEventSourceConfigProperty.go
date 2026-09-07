package lambda


// Experimental.
type AwsEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#consumer_group_id AwsEventSourceMapping#consumer_group_id}.
	// Experimental.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// schema_registry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#schema_registry_config AwsEventSourceMapping#schema_registry_config}
	// Experimental.
	SchemaRegistryConfig *AwsEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty `field:"optional" json:"schemaRegistryConfig" yaml:"schemaRegistryConfig"`
}

