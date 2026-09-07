package kinesisfirehose


// Experimental.
type AwsDeliveryStream_OutputFormatConfigurationProperty struct {
	// serializer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#serializer AwsDeliveryStream#serializer}
	// Experimental.
	Serializer *AwsDeliveryStream_SerializerProperty `field:"required" json:"serializer" yaml:"serializer"`
}

