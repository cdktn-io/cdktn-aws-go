package awskinesisfirehose


// Experimental.
type TfDeliveryStream_ServerSideEncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enabled TfDeliveryStream#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#key_arn TfDeliveryStream#key_arn}.
	// Experimental.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#key_type TfDeliveryStream#key_type}.
	// Experimental.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

