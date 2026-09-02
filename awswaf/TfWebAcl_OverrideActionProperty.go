package awswaf


// Experimental.
type TfWebAcl_OverrideActionProperty struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#count TfWebAcl#count}
	// Experimental.
	Count *TfWebAcl_RuleOverrideActionCountProperty `field:"optional" json:"count" yaml:"count"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#none TfWebAcl#none}
	// Experimental.
	None *TfWebAcl_NoneProperty `field:"optional" json:"none" yaml:"none"`
}

