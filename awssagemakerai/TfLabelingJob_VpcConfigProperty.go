package awssagemakerai


// Experimental.
type TfLabelingJob_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#security_group_ids TfLabelingJob#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#subnets TfLabelingJob#subnets}.
	// Experimental.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

