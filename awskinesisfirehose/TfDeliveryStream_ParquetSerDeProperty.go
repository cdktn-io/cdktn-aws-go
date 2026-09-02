package awskinesisfirehose


// Experimental.
type TfDeliveryStream_ParquetSerDeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#block_size_bytes TfDeliveryStream#block_size_bytes}.
	// Experimental.
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#compression TfDeliveryStream#compression}.
	// Experimental.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enable_dictionary_compression TfDeliveryStream#enable_dictionary_compression}.
	// Experimental.
	EnableDictionaryCompression interface{} `field:"optional" json:"enableDictionaryCompression" yaml:"enableDictionaryCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#max_padding_bytes TfDeliveryStream#max_padding_bytes}.
	// Experimental.
	MaxPaddingBytes *float64 `field:"optional" json:"maxPaddingBytes" yaml:"maxPaddingBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#page_size_bytes TfDeliveryStream#page_size_bytes}.
	// Experimental.
	PageSizeBytes *float64 `field:"optional" json:"pageSizeBytes" yaml:"pageSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#writer_version TfDeliveryStream#writer_version}.
	// Experimental.
	WriterVersion *string `field:"optional" json:"writerVersion" yaml:"writerVersion"`
}

