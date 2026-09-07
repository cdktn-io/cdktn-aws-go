package opensearch


// Experimental.
type AwsDomain_IdentityCenterOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#enabled_api_access AwsDomain#enabled_api_access}.
	// Experimental.
	EnabledApiAccess interface{} `field:"optional" json:"enabledApiAccess" yaml:"enabledApiAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#identity_center_instance_arn AwsDomain#identity_center_instance_arn}.
	// Experimental.
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#roles_key AwsDomain#roles_key}.
	// Experimental.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#subject_key AwsDomain#subject_key}.
	// Experimental.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

