package awsappmesh


// Experimental.
type TfVirtualGateway_SpecListenerTlsValidationSubjectAlternativeNamesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#match TfVirtualGateway#match}
	// Experimental.
	Match *TfVirtualGateway_SpecListenerTlsValidationSubjectAlternativeNamesMatchProperty `field:"required" json:"match" yaml:"match"`
}

