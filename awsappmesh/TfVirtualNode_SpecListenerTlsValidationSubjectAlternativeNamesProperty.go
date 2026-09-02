package awsappmesh


// Experimental.
type TfVirtualNode_SpecListenerTlsValidationSubjectAlternativeNamesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#match TfVirtualNode#match}
	// Experimental.
	Match *TfVirtualNode_SpecListenerTlsValidationSubjectAlternativeNamesMatchProperty `field:"required" json:"match" yaml:"match"`
}

