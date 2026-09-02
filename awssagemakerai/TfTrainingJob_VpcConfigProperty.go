package awssagemakerai


// Experimental.
type TfTrainingJob_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#security_group_ids TfTrainingJob#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#subnets TfTrainingJob#subnets}.
	// Experimental.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

