package awscodebuild


// Experimental.
type TfProject_BuildBatchConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#service_role TfProject#service_role}.
	// Experimental.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#combine_artifacts TfProject#combine_artifacts}.
	// Experimental.
	CombineArtifacts interface{} `field:"optional" json:"combineArtifacts" yaml:"combineArtifacts"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#restrictions TfProject#restrictions}
	// Experimental.
	Restrictions *TfProject_RestrictionsProperty `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#timeout_in_mins TfProject#timeout_in_mins}.
	// Experimental.
	TimeoutInMins *float64 `field:"optional" json:"timeoutInMins" yaml:"timeoutInMins"`
}

