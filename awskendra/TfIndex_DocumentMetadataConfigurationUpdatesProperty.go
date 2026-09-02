package awskendra


// Experimental.
type TfIndex_DocumentMetadataConfigurationUpdatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#name TfIndex#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#type TfIndex#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// relevance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#relevance TfIndex#relevance}
	// Experimental.
	Relevance *TfIndex_RelevanceProperty `field:"optional" json:"relevance" yaml:"relevance"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#search TfIndex#search}
	// Experimental.
	Search *TfIndex_SearchProperty `field:"optional" json:"search" yaml:"search"`
}

