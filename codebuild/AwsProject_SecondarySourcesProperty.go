package codebuild


// Experimental.
type AwsProject_SecondarySourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#source_identifier AwsProject#source_identifier}.
	// Experimental.
	SourceIdentifier *string `field:"required" json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#type AwsProject#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#auth AwsProject#auth}
	// Experimental.
	Auth *AwsProject_SecondarySourcesAuthProperty `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#buildspec AwsProject#buildspec}.
	// Experimental.
	Buildspec *string `field:"optional" json:"buildspec" yaml:"buildspec"`
	// build_status_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#build_status_config AwsProject#build_status_config}
	// Experimental.
	BuildStatusConfig *AwsProject_SecondarySourcesBuildStatusConfigProperty `field:"optional" json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#git_clone_depth AwsProject#git_clone_depth}.
	// Experimental.
	GitCloneDepth *float64 `field:"optional" json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// git_submodules_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#git_submodules_config AwsProject#git_submodules_config}
	// Experimental.
	GitSubmodulesConfig *AwsProject_SecondarySourcesGitSubmodulesConfigProperty `field:"optional" json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#insecure_ssl AwsProject#insecure_ssl}.
	// Experimental.
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#location AwsProject#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#report_build_status AwsProject#report_build_status}.
	// Experimental.
	ReportBuildStatus interface{} `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
}

