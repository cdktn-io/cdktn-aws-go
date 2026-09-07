package workspaces


// Experimental.
type AwsIpGroup_RulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_ip_group#source AwsIpGroup#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_ip_group#description AwsIpGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

