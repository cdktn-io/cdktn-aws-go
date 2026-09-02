package awsce


// Experimental.
type DataTfTags_NotProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#cost_category DataTfTags#cost_category}
	// Experimental.
	CostCategory *DataTfTags_FilterNotCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#dimension DataTfTags#dimension}
	// Experimental.
	Dimension *DataTfTags_FilterNotDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#tags DataTfTags#tags}
	// Experimental.
	Tags *DataTfTags_FilterNotTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

