package awsquicksight


// Experimental.
type AwsQuicksightDataSource_TwitterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#max_rows AwsQuicksightDataSource#max_rows}.
	// Experimental.
	MaxRows *float64 `field:"required" json:"maxRows" yaml:"maxRows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#query AwsQuicksightDataSource#query}.
	// Experimental.
	Query *string `field:"required" json:"query" yaml:"query"`
}

