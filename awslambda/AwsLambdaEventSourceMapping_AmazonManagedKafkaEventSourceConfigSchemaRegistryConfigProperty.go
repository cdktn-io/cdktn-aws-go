package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#access_config AwsLambdaEventSourceMapping#access_config}
	// Experimental.
	AccessConfig interface{} `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#event_record_format AwsLambdaEventSourceMapping#event_record_format}.
	// Experimental.
	EventRecordFormat *string `field:"optional" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#schema_registry_uri AwsLambdaEventSourceMapping#schema_registry_uri}.
	// Experimental.
	SchemaRegistryUri *string `field:"optional" json:"schemaRegistryUri" yaml:"schemaRegistryUri"`
	// schema_validation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#schema_validation_config AwsLambdaEventSourceMapping#schema_validation_config}
	// Experimental.
	SchemaValidationConfig interface{} `field:"optional" json:"schemaValidationConfig" yaml:"schemaValidationConfig"`
}

