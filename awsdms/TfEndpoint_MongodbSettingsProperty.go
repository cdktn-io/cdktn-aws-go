package awsdms


// Experimental.
type TfEndpoint_MongodbSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_mechanism TfEndpoint#auth_mechanism}.
	// Experimental.
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_source TfEndpoint#auth_source}.
	// Experimental.
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_type TfEndpoint#auth_type}.
	// Experimental.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#docs_to_investigate TfEndpoint#docs_to_investigate}.
	// Experimental.
	DocsToInvestigate *string `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#extract_doc_id TfEndpoint#extract_doc_id}.
	// Experimental.
	ExtractDocId *string `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#nesting_level TfEndpoint#nesting_level}.
	// Experimental.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_update_lookup TfEndpoint#use_update_lookup}.
	// Experimental.
	UseUpdateLookup interface{} `field:"optional" json:"useUpdateLookup" yaml:"useUpdateLookup"`
}

