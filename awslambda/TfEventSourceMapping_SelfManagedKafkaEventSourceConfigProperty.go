package awslambda


// Experimental.
type TfEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#consumer_group_id TfEventSourceMapping#consumer_group_id}.
	// Experimental.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// schema_registry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#schema_registry_config TfEventSourceMapping#schema_registry_config}
	// Experimental.
	SchemaRegistryConfig *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty `field:"optional" json:"schemaRegistryConfig" yaml:"schemaRegistryConfig"`
}

