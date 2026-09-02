package awscodebuild


// Experimental.
type TfProject_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#type TfProject#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#auth TfProject#auth}
	// Experimental.
	Auth *TfProject_SourceAuthProperty `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#buildspec TfProject#buildspec}.
	// Experimental.
	Buildspec *string `field:"optional" json:"buildspec" yaml:"buildspec"`
	// build_status_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#build_status_config TfProject#build_status_config}
	// Experimental.
	BuildStatusConfig *TfProject_SourceBuildStatusConfigProperty `field:"optional" json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#git_clone_depth TfProject#git_clone_depth}.
	// Experimental.
	GitCloneDepth *float64 `field:"optional" json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// git_submodules_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#git_submodules_config TfProject#git_submodules_config}
	// Experimental.
	GitSubmodulesConfig *TfProject_SourceGitSubmodulesConfigProperty `field:"optional" json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#insecure_ssl TfProject#insecure_ssl}.
	// Experimental.
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#location TfProject#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#report_build_status TfProject#report_build_status}.
	// Experimental.
	ReportBuildStatus interface{} `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
}

