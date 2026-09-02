package awsdynamodb


// Experimental.
type TfTableExport_IncrementalExportSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_from_time TfTableExport#export_from_time}.
	// Experimental.
	ExportFromTime *string `field:"optional" json:"exportFromTime" yaml:"exportFromTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_to_time TfTableExport#export_to_time}.
	// Experimental.
	ExportToTime *string `field:"optional" json:"exportToTime" yaml:"exportToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_view_type TfTableExport#export_view_type}.
	// Experimental.
	ExportViewType *string `field:"optional" json:"exportViewType" yaml:"exportViewType"`
}

