package awsquicksight


// Experimental.
type TfDashboard_SourceTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#arn TfDashboard#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// data_set_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_set_references TfDashboard#data_set_references}
	// Experimental.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

