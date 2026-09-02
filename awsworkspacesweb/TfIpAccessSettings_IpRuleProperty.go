package awsworkspacesweb


// Experimental.
type TfIpAccessSettings_IpRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_ip_access_settings#ip_range TfIpAccessSettings#ip_range}.
	// Experimental.
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_ip_access_settings#description TfIpAccessSettings#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

