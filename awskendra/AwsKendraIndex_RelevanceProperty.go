package awskendra


// Experimental.
type AwsKendraIndex_RelevanceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#duration AwsKendraIndex#duration}.
	// Experimental.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#freshness AwsKendraIndex#freshness}.
	// Experimental.
	Freshness interface{} `field:"optional" json:"freshness" yaml:"freshness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#importance AwsKendraIndex#importance}.
	// Experimental.
	Importance *float64 `field:"optional" json:"importance" yaml:"importance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#rank_order AwsKendraIndex#rank_order}.
	// Experimental.
	RankOrder *string `field:"optional" json:"rankOrder" yaml:"rankOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#values_importance_map AwsKendraIndex#values_importance_map}.
	// Experimental.
	ValuesImportanceMap *map[string]*float64 `field:"optional" json:"valuesImportanceMap" yaml:"valuesImportanceMap"`
}

