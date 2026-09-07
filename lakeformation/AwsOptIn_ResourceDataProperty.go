package lakeformation


// Experimental.
type AwsOptIn_ResourceDataProperty struct {
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#catalog AwsOptIn#catalog}
	// Experimental.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#database AwsOptIn#database}
	// Experimental.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// data_cells_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#data_cells_filter AwsOptIn#data_cells_filter}
	// Experimental.
	DataCellsFilter interface{} `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// data_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#data_location AwsOptIn#data_location}
	// Experimental.
	DataLocation interface{} `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// lf_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#lf_tag AwsOptIn#lf_tag}
	// Experimental.
	LfTag interface{} `field:"optional" json:"lfTag" yaml:"lfTag"`
	// lf_tag_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#lf_tag_expression AwsOptIn#lf_tag_expression}
	// Experimental.
	LfTagExpression interface{} `field:"optional" json:"lfTagExpression" yaml:"lfTagExpression"`
	// lf_tag_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#lf_tag_policy AwsOptIn#lf_tag_policy}
	// Experimental.
	LfTagPolicy interface{} `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#table AwsOptIn#table}
	// Experimental.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#table_with_columns AwsOptIn#table_with_columns}
	// Experimental.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

