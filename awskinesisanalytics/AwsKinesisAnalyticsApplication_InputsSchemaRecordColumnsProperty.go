package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_InputsSchemaRecordColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name AwsKinesisAnalyticsApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#sql_type AwsKinesisAnalyticsApplication#sql_type}.
	// Experimental.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#mapping AwsKinesisAnalyticsApplication#mapping}.
	// Experimental.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

