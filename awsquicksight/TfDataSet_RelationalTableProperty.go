package awsquicksight


// Experimental.
type TfDataSet_RelationalTableProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_source_arn TfDataSet#data_source_arn}.
	// Experimental.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// input_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#input_columns TfDataSet#input_columns}
	// Experimental.
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#name TfDataSet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#catalog TfDataSet#catalog}.
	// Experimental.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#schema TfDataSet#schema}.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

