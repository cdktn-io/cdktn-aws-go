package awskinesisfirehose


// Experimental.
type TfDeliveryStream_RequestConfigurationProperty struct {
	// common_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#common_attributes TfDeliveryStream#common_attributes}
	// Experimental.
	CommonAttributes interface{} `field:"optional" json:"commonAttributes" yaml:"commonAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#content_encoding TfDeliveryStream#content_encoding}.
	// Experimental.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
}

