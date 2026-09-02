package awsquicksight


// Experimental.
type TfDashboard_DataSetReferencesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_set_arn TfDashboard#data_set_arn}.
	// Experimental.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_set_placeholder TfDashboard#data_set_placeholder}.
	// Experimental.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

