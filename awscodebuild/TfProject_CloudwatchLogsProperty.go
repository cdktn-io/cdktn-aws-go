package awscodebuild


// Experimental.
type TfProject_CloudwatchLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#group_name TfProject#group_name}.
	// Experimental.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#status TfProject#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#stream_name TfProject#stream_name}.
	// Experimental.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

