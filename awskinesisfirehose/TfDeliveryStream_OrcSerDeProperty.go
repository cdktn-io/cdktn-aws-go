package awskinesisfirehose


// Experimental.
type TfDeliveryStream_OrcSerDeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#block_size_bytes TfDeliveryStream#block_size_bytes}.
	// Experimental.
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#bloom_filter_columns TfDeliveryStream#bloom_filter_columns}.
	// Experimental.
	BloomFilterColumns *[]*string `field:"optional" json:"bloomFilterColumns" yaml:"bloomFilterColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#bloom_filter_false_positive_probability TfDeliveryStream#bloom_filter_false_positive_probability}.
	// Experimental.
	BloomFilterFalsePositiveProbability *float64 `field:"optional" json:"bloomFilterFalsePositiveProbability" yaml:"bloomFilterFalsePositiveProbability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#compression TfDeliveryStream#compression}.
	// Experimental.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#dictionary_key_threshold TfDeliveryStream#dictionary_key_threshold}.
	// Experimental.
	DictionaryKeyThreshold *float64 `field:"optional" json:"dictionaryKeyThreshold" yaml:"dictionaryKeyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enable_padding TfDeliveryStream#enable_padding}.
	// Experimental.
	EnablePadding interface{} `field:"optional" json:"enablePadding" yaml:"enablePadding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#format_version TfDeliveryStream#format_version}.
	// Experimental.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#padding_tolerance TfDeliveryStream#padding_tolerance}.
	// Experimental.
	PaddingTolerance *float64 `field:"optional" json:"paddingTolerance" yaml:"paddingTolerance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#row_index_stride TfDeliveryStream#row_index_stride}.
	// Experimental.
	RowIndexStride *float64 `field:"optional" json:"rowIndexStride" yaml:"rowIndexStride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#stripe_size_bytes TfDeliveryStream#stripe_size_bytes}.
	// Experimental.
	StripeSizeBytes *float64 `field:"optional" json:"stripeSizeBytes" yaml:"stripeSizeBytes"`
}

