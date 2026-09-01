package awswaf


// Experimental.
type AwsWafv2WebAcl_DefaultActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#allow AwsWafv2WebAcl#allow}
	// Experimental.
	Allow *AwsWafv2WebAcl_DefaultActionAllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#block AwsWafv2WebAcl#block}
	// Experimental.
	Block *AwsWafv2WebAcl_DefaultActionBlockProperty `field:"optional" json:"block" yaml:"block"`
}

