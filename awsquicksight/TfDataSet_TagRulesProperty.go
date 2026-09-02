package awsquicksight


// Experimental.
type TfDataSet_TagRulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name TfDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_key TfDataSet#tag_key}.
	// Experimental.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#match_all_value TfDataSet#match_all_value}.
	// Experimental.
	MatchAllValue *string `field:"optional" json:"matchAllValue" yaml:"matchAllValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_multi_value_delimiter TfDataSet#tag_multi_value_delimiter}.
	// Experimental.
	TagMultiValueDelimiter *string `field:"optional" json:"tagMultiValueDelimiter" yaml:"tagMultiValueDelimiter"`
}

