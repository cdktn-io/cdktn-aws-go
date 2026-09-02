package awskinesisfirehose


// Experimental.
type TfDeliveryStream_InputFormatConfigurationProperty struct {
	// deserializer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#deserializer TfDeliveryStream#deserializer}
	// Experimental.
	Deserializer *TfDeliveryStream_DeserializerProperty `field:"required" json:"deserializer" yaml:"deserializer"`
}

