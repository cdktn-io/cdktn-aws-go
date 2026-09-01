package awskendra


// Experimental.
type AwsKendraIndex_DocumentMetadataConfigurationUpdatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#name AwsKendraIndex#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#type AwsKendraIndex#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// relevance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#relevance AwsKendraIndex#relevance}
	// Experimental.
	Relevance *AwsKendraIndex_RelevanceProperty `field:"optional" json:"relevance" yaml:"relevance"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#search AwsKendraIndex#search}
	// Experimental.
	Search *AwsKendraIndex_SearchProperty `field:"optional" json:"search" yaml:"search"`
}

