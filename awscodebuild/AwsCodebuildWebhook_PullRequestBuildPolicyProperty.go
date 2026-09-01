package awscodebuild


// Experimental.
type AwsCodebuildWebhook_PullRequestBuildPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#requires_comment_approval AwsCodebuildWebhook#requires_comment_approval}.
	// Experimental.
	RequiresCommentApproval *string `field:"required" json:"requiresCommentApproval" yaml:"requiresCommentApproval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#approver_roles AwsCodebuildWebhook#approver_roles}.
	// Experimental.
	ApproverRoles *[]*string `field:"optional" json:"approverRoles" yaml:"approverRoles"`
}

