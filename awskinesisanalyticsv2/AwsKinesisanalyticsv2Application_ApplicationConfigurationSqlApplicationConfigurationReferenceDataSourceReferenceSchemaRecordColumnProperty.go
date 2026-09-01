package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name AwsKinesisanalyticsv2Application#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#sql_type AwsKinesisanalyticsv2Application#sql_type}.
	// Experimental.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#mapping AwsKinesisanalyticsv2Application#mapping}.
	// Experimental.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

