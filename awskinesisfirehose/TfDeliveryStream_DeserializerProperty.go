package awskinesisfirehose


// Experimental.
type TfDeliveryStream_DeserializerProperty struct {
	// hive_json_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hive_json_ser_de TfDeliveryStream#hive_json_ser_de}
	// Experimental.
	HiveJsonSerDe *TfDeliveryStream_HiveJsonSerDeProperty `field:"optional" json:"hiveJsonSerDe" yaml:"hiveJsonSerDe"`
	// open_x_json_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#open_x_json_ser_de TfDeliveryStream#open_x_json_ser_de}
	// Experimental.
	OpenXJsonSerDe *TfDeliveryStream_OpenXJsonSerDeProperty `field:"optional" json:"openXJsonSerDe" yaml:"openXJsonSerDe"`
}

