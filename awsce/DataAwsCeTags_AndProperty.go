package awsce


// Experimental.
type DataAwsCeTags_AndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#cost_category DataAwsCeTags#cost_category}
	// Experimental.
	CostCategory *DataAwsCeTags_FilterAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#dimension DataAwsCeTags#dimension}
	// Experimental.
	Dimension *DataAwsCeTags_FilterAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#tags DataAwsCeTags#tags}
	// Experimental.
	Tags *DataAwsCeTags_FilterAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

