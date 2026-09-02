package awskendra


// Experimental.
type TfIndex_RelevanceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#duration TfIndex#duration}.
	// Experimental.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#freshness TfIndex#freshness}.
	// Experimental.
	Freshness interface{} `field:"optional" json:"freshness" yaml:"freshness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#importance TfIndex#importance}.
	// Experimental.
	Importance *float64 `field:"optional" json:"importance" yaml:"importance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#rank_order TfIndex#rank_order}.
	// Experimental.
	RankOrder *string `field:"optional" json:"rankOrder" yaml:"rankOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#values_importance_map TfIndex#values_importance_map}.
	// Experimental.
	ValuesImportanceMap *map[string]*float64 `field:"optional" json:"valuesImportanceMap" yaml:"valuesImportanceMap"`
}

