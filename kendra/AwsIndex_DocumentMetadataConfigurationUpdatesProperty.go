package kendra


// Experimental.
type AwsIndex_DocumentMetadataConfigurationUpdatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#name AwsIndex#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#type AwsIndex#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// relevance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#relevance AwsIndex#relevance}
	// Experimental.
	Relevance *AwsIndex_RelevanceProperty `field:"optional" json:"relevance" yaml:"relevance"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#search AwsIndex#search}
	// Experimental.
	Search *AwsIndex_SearchProperty `field:"optional" json:"search" yaml:"search"`
}

