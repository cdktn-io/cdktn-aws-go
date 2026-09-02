package awskinesisfirehose


// Experimental.
type TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enabled TfDeliveryStream#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// processors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processors TfDeliveryStream#processors}
	// Experimental.
	Processors interface{} `field:"optional" json:"processors" yaml:"processors"`
}

