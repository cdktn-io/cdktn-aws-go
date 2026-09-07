package qldb


// Experimental.
type AwsStream_KinesisConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qldb_stream#stream_arn AwsStream#stream_arn}.
	// Experimental.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qldb_stream#aggregation_enabled AwsStream#aggregation_enabled}.
	// Experimental.
	AggregationEnabled interface{} `field:"optional" json:"aggregationEnabled" yaml:"aggregationEnabled"`
}

