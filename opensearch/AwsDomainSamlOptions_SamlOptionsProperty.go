package opensearch


// Experimental.
type AwsDomainSamlOptions_SamlOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#enabled AwsDomainSamlOptions#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// idp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#idp AwsDomainSamlOptions#idp}
	// Experimental.
	Idp *AwsDomainSamlOptions_IdpProperty `field:"optional" json:"idp" yaml:"idp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#master_backend_role AwsDomainSamlOptions#master_backend_role}.
	// Experimental.
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#master_user_name AwsDomainSamlOptions#master_user_name}.
	// Experimental.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#roles_key AwsDomainSamlOptions#roles_key}.
	// Experimental.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#session_timeout_minutes AwsDomainSamlOptions#session_timeout_minutes}.
	// Experimental.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#subject_key AwsDomainSamlOptions#subject_key}.
	// Experimental.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

