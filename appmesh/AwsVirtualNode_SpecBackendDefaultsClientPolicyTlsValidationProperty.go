package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#trust AwsVirtualNode#trust}
	// Experimental.
	Trust *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#subject_alternative_names AwsVirtualNode#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

