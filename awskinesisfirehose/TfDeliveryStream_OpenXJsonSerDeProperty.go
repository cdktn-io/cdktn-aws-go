package awskinesisfirehose


// Experimental.
type TfDeliveryStream_OpenXJsonSerDeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#case_insensitive TfDeliveryStream#case_insensitive}.
	// Experimental.
	CaseInsensitive interface{} `field:"optional" json:"caseInsensitive" yaml:"caseInsensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#column_to_json_key_mappings TfDeliveryStream#column_to_json_key_mappings}.
	// Experimental.
	ColumnToJsonKeyMappings *map[string]*string `field:"optional" json:"columnToJsonKeyMappings" yaml:"columnToJsonKeyMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#convert_dots_in_json_keys_to_underscores TfDeliveryStream#convert_dots_in_json_keys_to_underscores}.
	// Experimental.
	ConvertDotsInJsonKeysToUnderscores interface{} `field:"optional" json:"convertDotsInJsonKeysToUnderscores" yaml:"convertDotsInJsonKeysToUnderscores"`
}

