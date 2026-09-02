package awswaf


// Experimental.
type TfWebAcl_DefaultActionProperty struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#allow TfWebAcl#allow}
	// Experimental.
	Allow *TfWebAcl_DefaultActionAllowProperty `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#block TfWebAcl#block}
	// Experimental.
	Block *TfWebAcl_DefaultActionBlockProperty `field:"optional" json:"block" yaml:"block"`
}

