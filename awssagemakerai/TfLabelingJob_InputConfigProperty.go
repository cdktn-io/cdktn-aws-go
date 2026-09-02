package awssagemakerai


// Experimental.
type TfLabelingJob_InputConfigProperty struct {
	// data_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#data_attributes TfLabelingJob#data_attributes}
	// Experimental.
	DataAttributes interface{} `field:"optional" json:"dataAttributes" yaml:"dataAttributes"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#data_source TfLabelingJob#data_source}
	// Experimental.
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
}

