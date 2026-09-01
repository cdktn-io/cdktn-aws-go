package awsquicksight


// Experimental.
type AwsQuicksightDataSet_RowLevelPermissionTagConfigurationProperty struct {
	// tag_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_rules AwsQuicksightDataSet#tag_rules}
	// Experimental.
	TagRules interface{} `field:"required" json:"tagRules" yaml:"tagRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#status AwsQuicksightDataSet#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

