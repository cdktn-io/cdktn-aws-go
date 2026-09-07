package dynamodb


// Experimental.
type AwsTableExport_IncrementalExportSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_from_time AwsTableExport#export_from_time}.
	// Experimental.
	ExportFromTime *string `field:"optional" json:"exportFromTime" yaml:"exportFromTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_to_time AwsTableExport#export_to_time}.
	// Experimental.
	ExportToTime *string `field:"optional" json:"exportToTime" yaml:"exportToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_view_type AwsTableExport#export_view_type}.
	// Experimental.
	ExportViewType *string `field:"optional" json:"exportViewType" yaml:"exportViewType"`
}

