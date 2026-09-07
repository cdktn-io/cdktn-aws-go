package s3tables


// Experimental.
type AwsTableBucket_SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#non_current_days AwsTableBucket#non_current_days}.
	// Experimental.
	NonCurrentDays *float64 `field:"optional" json:"nonCurrentDays" yaml:"nonCurrentDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#unreferenced_days AwsTableBucket#unreferenced_days}.
	// Experimental.
	UnreferencedDays *float64 `field:"optional" json:"unreferencedDays" yaml:"unreferencedDays"`
}

