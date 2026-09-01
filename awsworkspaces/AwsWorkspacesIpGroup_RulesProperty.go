package awsworkspaces


// Experimental.
type AwsWorkspacesIpGroup_RulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_ip_group#source AwsWorkspacesIpGroup#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_ip_group#description AwsWorkspacesIpGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

