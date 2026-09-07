package kinesisfirehose


// Experimental.
type AwsDeliveryStream_DeserializerProperty struct {
	// hive_json_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#hive_json_ser_de AwsDeliveryStream#hive_json_ser_de}
	// Experimental.
	HiveJsonSerDe *AwsDeliveryStream_HiveJsonSerDeProperty `field:"optional" json:"hiveJsonSerDe" yaml:"hiveJsonSerDe"`
	// open_x_json_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#open_x_json_ser_de AwsDeliveryStream#open_x_json_ser_de}
	// Experimental.
	OpenXJsonSerDe *AwsDeliveryStream_OpenXJsonSerDeProperty `field:"optional" json:"openXJsonSerDe" yaml:"openXJsonSerDe"`
}

