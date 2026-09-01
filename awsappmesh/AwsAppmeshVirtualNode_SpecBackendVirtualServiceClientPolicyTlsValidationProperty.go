package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#trust AwsAppmeshVirtualNode#trust}
	// Experimental.
	Trust *AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#subject_alternative_names AwsAppmeshVirtualNode#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

