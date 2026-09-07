package waf


// Experimental.
type AwsWebAcl_DefaultActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#allow AwsWebAcl#allow}
	// Experimental.
	Allow *AwsWebAcl_DefaultActionAllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#block AwsWebAcl#block}
	// Experimental.
	Block *AwsWebAcl_DefaultActionBlockProperty `field:"optional" json:"block" yaml:"block"`
}

