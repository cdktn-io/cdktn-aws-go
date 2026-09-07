package kinesisfirehose


// Experimental.
type AwsDeliveryStream_InputFormatConfigurationProperty struct {
	// deserializer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#deserializer AwsDeliveryStream#deserializer}
	// Experimental.
	Deserializer *AwsDeliveryStream_DeserializerProperty `field:"required" json:"deserializer" yaml:"deserializer"`
}

