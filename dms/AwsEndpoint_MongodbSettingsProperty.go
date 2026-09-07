package dms


// Experimental.
type AwsEndpoint_MongodbSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_mechanism AwsEndpoint#auth_mechanism}.
	// Experimental.
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_source AwsEndpoint#auth_source}.
	// Experimental.
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_type AwsEndpoint#auth_type}.
	// Experimental.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#docs_to_investigate AwsEndpoint#docs_to_investigate}.
	// Experimental.
	DocsToInvestigate *string `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#extract_doc_id AwsEndpoint#extract_doc_id}.
	// Experimental.
	ExtractDocId *string `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#nesting_level AwsEndpoint#nesting_level}.
	// Experimental.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_update_lookup AwsEndpoint#use_update_lookup}.
	// Experimental.
	UseUpdateLookup interface{} `field:"optional" json:"useUpdateLookup" yaml:"useUpdateLookup"`
}

