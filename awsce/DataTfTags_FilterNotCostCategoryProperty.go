package awsce


// Experimental.
type DataTfTags_FilterNotCostCategoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#key DataTfTags#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#match_options DataTfTags#match_options}.
	// Experimental.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ce_tags#values DataTfTags#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

