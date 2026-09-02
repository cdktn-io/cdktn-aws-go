package awskinesisfirehose


// Experimental.
type TfDeliveryStream_SerializerProperty struct {
	// orc_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#orc_ser_de TfDeliveryStream#orc_ser_de}
	// Experimental.
	OrcSerDe *TfDeliveryStream_OrcSerDeProperty `field:"optional" json:"orcSerDe" yaml:"orcSerDe"`
	// parquet_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#parquet_ser_de TfDeliveryStream#parquet_ser_de}
	// Experimental.
	ParquetSerDe *TfDeliveryStream_ParquetSerDeProperty `field:"optional" json:"parquetSerDe" yaml:"parquetSerDe"`
}

