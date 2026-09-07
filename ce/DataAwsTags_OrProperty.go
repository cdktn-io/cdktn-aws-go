package ce


// Experimental.
type DataAwsTags_OrProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#cost_category DataAwsTags#cost_category}
	// Experimental.
	CostCategory *DataAwsTags_FilterOrCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#dimension DataAwsTags#dimension}
	// Experimental.
	Dimension *DataAwsTags_FilterOrDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#tags DataAwsTags#tags}
	// Experimental.
	Tags *DataAwsTags_FilterOrTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

