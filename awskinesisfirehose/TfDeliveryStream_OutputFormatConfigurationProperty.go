package awskinesisfirehose


// Experimental.
type TfDeliveryStream_OutputFormatConfigurationProperty struct {
	// serializer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#serializer TfDeliveryStream#serializer}
	// Experimental.
	Serializer *TfDeliveryStream_SerializerProperty `field:"required" json:"serializer" yaml:"serializer"`
}

