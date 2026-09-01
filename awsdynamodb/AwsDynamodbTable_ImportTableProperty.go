package awsdynamodb


// Experimental.
type AwsDynamodbTable_ImportTableProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#input_format AwsDynamodbTable#input_format}.
	// Experimental.
	InputFormat *string `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// s3_bucket_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#s3_bucket_source AwsDynamodbTable#s3_bucket_source}
	// Experimental.
	S3BucketSource *AwsDynamodbTable_S3BucketSourceProperty `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#input_compression_type AwsDynamodbTable#input_compression_type}.
	// Experimental.
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// input_format_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#input_format_options AwsDynamodbTable#input_format_options}
	// Experimental.
	InputFormatOptions *AwsDynamodbTable_InputFormatOptionsProperty `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
}

