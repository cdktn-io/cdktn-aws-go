package awselasticsearch


// Experimental.
type AwsElasticsearchDomainSamlOptions_IdpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#entity_id AwsElasticsearchDomainSamlOptions#entity_id}.
	// Experimental.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain_saml_options#metadata_content AwsElasticsearchDomainSamlOptions#metadata_content}.
	// Experimental.
	MetadataContent *string `field:"required" json:"metadataContent" yaml:"metadataContent"`
}

