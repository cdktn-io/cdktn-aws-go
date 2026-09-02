package awsresiliencehubv2


// Experimental.
type TfInputSource_ResourceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#cfn_stack_arn TfInputSource#cfn_stack_arn}.
	// Experimental.
	CfnStackArn *string `field:"optional" json:"cfnStackArn" yaml:"cfnStackArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#design_file_s3_url TfInputSource#design_file_s3_url}.
	// Experimental.
	DesignFileS3Url *string `field:"optional" json:"designFileS3Url" yaml:"designFileS3Url"`
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#eks TfInputSource#eks}
	// Experimental.
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
	// resource_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#resource_tag TfInputSource#resource_tag}
	// Experimental.
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#tf_state_file_url TfInputSource#tf_state_file_url}.
	// Experimental.
	TfStateFileUrl *string `field:"optional" json:"tfStateFileUrl" yaml:"tfStateFileUrl"`
}

