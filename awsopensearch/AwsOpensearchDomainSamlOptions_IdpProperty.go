package awsopensearch


// Experimental.
type AwsOpensearchDomainSamlOptions_IdpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#entity_id AwsOpensearchDomainSamlOptions#entity_id}.
	// Experimental.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_saml_options#metadata_content AwsOpensearchDomainSamlOptions#metadata_content}.
	// Experimental.
	MetadataContent *string `field:"required" json:"metadataContent" yaml:"metadataContent"`
}

