package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerTlsValidationSubjectAlternativeNamesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#match AwsAppmeshVirtualNode#match}
	// Experimental.
	Match *AwsAppmeshVirtualNode_SpecListenerTlsValidationSubjectAlternativeNamesMatchProperty `field:"required" json:"match" yaml:"match"`
}

