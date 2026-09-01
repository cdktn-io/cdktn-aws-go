package awslambda


// Experimental.
type AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#database_name AwsLambdaEventSourceMapping#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#collection_name AwsLambdaEventSourceMapping#collection_name}.
	// Experimental.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#full_document AwsLambdaEventSourceMapping#full_document}.
	// Experimental.
	FullDocument *string `field:"optional" json:"fullDocument" yaml:"fullDocument"`
}

