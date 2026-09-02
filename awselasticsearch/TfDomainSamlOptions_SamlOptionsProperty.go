package awselasticsearch


// Experimental.
type TfDomainSamlOptions_SamlOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#enabled TfDomainSamlOptions#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// idp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#idp TfDomainSamlOptions#idp}
	// Experimental.
	Idp *TfDomainSamlOptions_IdpProperty `field:"optional" json:"idp" yaml:"idp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#master_backend_role TfDomainSamlOptions#master_backend_role}.
	// Experimental.
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#master_user_name TfDomainSamlOptions#master_user_name}.
	// Experimental.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#roles_key TfDomainSamlOptions#roles_key}.
	// Experimental.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#session_timeout_minutes TfDomainSamlOptions#session_timeout_minutes}.
	// Experimental.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#subject_key TfDomainSamlOptions#subject_key}.
	// Experimental.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

