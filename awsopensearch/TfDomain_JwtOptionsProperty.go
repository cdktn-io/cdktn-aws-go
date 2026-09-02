package awsopensearch


// Experimental.
type TfDomain_JwtOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#jwks_url TfDomain#jwks_url}.
	// Experimental.
	JwksUrl *string `field:"optional" json:"jwksUrl" yaml:"jwksUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#public_key TfDomain#public_key}.
	// Experimental.
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#roles_key TfDomain#roles_key}.
	// Experimental.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#subject_key TfDomain#subject_key}.
	// Experimental.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

