package awskinesisanalyticsv2


// Experimental.
type TfApplication_InputSchemaProperty struct {
	// record_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_column TfApplication#record_column}
	// Experimental.
	RecordColumn interface{} `field:"required" json:"recordColumn" yaml:"recordColumn"`
	// record_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_format TfApplication#record_format}
	// Experimental.
	RecordFormat *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_encoding TfApplication#record_encoding}.
	// Experimental.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

