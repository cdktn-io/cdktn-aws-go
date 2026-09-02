package awsopensearchserverless


// Experimental.
type TfCollectionGroup_CapacityLimitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#max_indexing_capacity_in_ocu TfCollectionGroup#max_indexing_capacity_in_ocu}.
	// Experimental.
	MaxIndexingCapacityInOcu *float64 `field:"optional" json:"maxIndexingCapacityInOcu" yaml:"maxIndexingCapacityInOcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#max_search_capacity_in_ocu TfCollectionGroup#max_search_capacity_in_ocu}.
	// Experimental.
	MaxSearchCapacityInOcu *float64 `field:"optional" json:"maxSearchCapacityInOcu" yaml:"maxSearchCapacityInOcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#min_indexing_capacity_in_ocu TfCollectionGroup#min_indexing_capacity_in_ocu}.
	// Experimental.
	MinIndexingCapacityInOcu *float64 `field:"optional" json:"minIndexingCapacityInOcu" yaml:"minIndexingCapacityInOcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#min_search_capacity_in_ocu TfCollectionGroup#min_search_capacity_in_ocu}.
	// Experimental.
	MinSearchCapacityInOcu *float64 `field:"optional" json:"minSearchCapacityInOcu" yaml:"minSearchCapacityInOcu"`
}

