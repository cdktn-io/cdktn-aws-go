package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendDefaultsClientPolicyTlsProperty struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#validation TfVirtualNode#validation}
	// Experimental.
	Validation *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationProperty `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate TfVirtualNode#certificate}
	// Experimental.
	Certificate *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateProperty `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#enforce TfVirtualNode#enforce}.
	// Experimental.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#ports TfVirtualNode#ports}.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

