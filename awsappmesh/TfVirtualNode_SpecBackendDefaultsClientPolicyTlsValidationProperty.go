package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#trust TfVirtualNode#trust}
	// Experimental.
	Trust *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#subject_alternative_names TfVirtualNode#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

