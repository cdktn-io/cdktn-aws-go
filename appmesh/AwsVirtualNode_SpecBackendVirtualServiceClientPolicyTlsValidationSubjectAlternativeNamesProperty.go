package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#match AwsVirtualNode#match}
	// Experimental.
	Match *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty `field:"required" json:"match" yaml:"match"`
}

