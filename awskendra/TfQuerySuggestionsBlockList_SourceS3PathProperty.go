package awskendra


// Experimental.
type TfQuerySuggestionsBlockList_SourceS3PathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_query_suggestions_block_list#bucket TfQuerySuggestionsBlockList#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_query_suggestions_block_list#key TfQuerySuggestionsBlockList#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

