package quicksight


// Experimental.
type AwsDataSet_ColumnLevelPermissionRulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_names AwsDataSet#column_names}.
	// Experimental.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#principals AwsDataSet#principals}.
	// Experimental.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

