package awsappmesh


// Experimental.
type TfVirtualNode_SpecListenerTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#trust TfVirtualNode#trust}
	// Experimental.
	Trust *TfVirtualNode_SpecListenerTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#subject_alternative_names TfVirtualNode#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *TfVirtualNode_SpecListenerTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

