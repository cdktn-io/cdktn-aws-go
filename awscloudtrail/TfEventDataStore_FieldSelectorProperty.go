package awscloudtrail


// Experimental.
type TfEventDataStore_FieldSelectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#ends_with TfEventDataStore#ends_with}.
	// Experimental.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#equals TfEventDataStore#equals}.
	// Experimental.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#field TfEventDataStore#field}.
	// Experimental.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#not_ends_with TfEventDataStore#not_ends_with}.
	// Experimental.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#not_equals TfEventDataStore#not_equals}.
	// Experimental.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#not_starts_with TfEventDataStore#not_starts_with}.
	// Experimental.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#starts_with TfEventDataStore#starts_with}.
	// Experimental.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

