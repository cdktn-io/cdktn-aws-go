package awskendra


// Experimental.
type TfDataSource_TargetDocumentAttributeValueProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#date_value TfDataSource#date_value}.
	// Experimental.
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#long_value TfDataSource#long_value}.
	// Experimental.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#string_list_value TfDataSource#string_list_value}.
	// Experimental.
	StringListValue *[]*string `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#string_value TfDataSource#string_value}.
	// Experimental.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

