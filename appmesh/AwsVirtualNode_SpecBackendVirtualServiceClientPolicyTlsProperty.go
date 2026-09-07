package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsProperty struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#validation AwsVirtualNode#validation}
	// Experimental.
	Validation *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate AwsVirtualNode#certificate}
	// Experimental.
	Certificate *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#enforce AwsVirtualNode#enforce}.
	// Experimental.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#ports AwsVirtualNode#ports}.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

