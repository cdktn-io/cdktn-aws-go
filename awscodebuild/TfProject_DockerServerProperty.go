package awscodebuild


// Experimental.
type TfProject_DockerServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#compute_type TfProject#compute_type}.
	// Experimental.
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#security_group_ids TfProject#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

