package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_InputSchemaProperty struct {
	// record_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_column AwsKinesisanalyticsv2Application#record_column}
	// Experimental.
	RecordColumn interface{} `field:"required" json:"recordColumn" yaml:"recordColumn"`
	// record_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_format AwsKinesisanalyticsv2Application#record_format}
	// Experimental.
	RecordFormat *AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_encoding AwsKinesisanalyticsv2Application#record_encoding}.
	// Experimental.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

