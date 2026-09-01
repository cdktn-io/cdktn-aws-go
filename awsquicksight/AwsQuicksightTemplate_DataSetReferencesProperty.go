package awsquicksight


// Experimental.
type AwsQuicksightTemplate_DataSetReferencesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#data_set_arn AwsQuicksightTemplate#data_set_arn}.
	// Experimental.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#data_set_placeholder AwsQuicksightTemplate#data_set_placeholder}.
	// Experimental.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

