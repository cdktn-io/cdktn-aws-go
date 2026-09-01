package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebIpAccessSettings_IpRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_ip_access_settings#ip_range AwsWorkspaceswebIpAccessSettings#ip_range}.
	// Experimental.
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_ip_access_settings#description AwsWorkspaceswebIpAccessSettings#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

