package awswaf


// Experimental.
type TfWebAcl_DefaultActionBlockProperty struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#custom_response TfWebAcl#custom_response}
	// Experimental.
	CustomResponse *TfWebAcl_DefaultActionBlockCustomResponseProperty `field:"optional" json:"customResponse" yaml:"customResponse"`
}

