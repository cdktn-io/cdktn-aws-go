package awsquicksight


// Experimental.
type TfDataSet_CustomSqlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_source_arn TfDataSet#data_source_arn}.
	// Experimental.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#name TfDataSet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#sql_query TfDataSet#sql_query}.
	// Experimental.
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#columns TfDataSet#columns}
	// Experimental.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

