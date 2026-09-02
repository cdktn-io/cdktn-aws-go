package awsglue


// Experimental.
type TfCatalogTable_ViewDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#definer TfCatalogTable#definer}.
	// Experimental.
	Definer *string `field:"optional" json:"definer" yaml:"definer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#is_protected TfCatalogTable#is_protected}.
	// Experimental.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#last_refresh_type TfCatalogTable#last_refresh_type}.
	// Experimental.
	LastRefreshType *string `field:"optional" json:"lastRefreshType" yaml:"lastRefreshType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#refresh_seconds TfCatalogTable#refresh_seconds}.
	// Experimental.
	RefreshSeconds *float64 `field:"optional" json:"refreshSeconds" yaml:"refreshSeconds"`
	// representations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#representations TfCatalogTable#representations}
	// Experimental.
	Representations interface{} `field:"optional" json:"representations" yaml:"representations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sub_objects TfCatalogTable#sub_objects}.
	// Experimental.
	SubObjects *[]*string `field:"optional" json:"subObjects" yaml:"subObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sub_object_version_ids TfCatalogTable#sub_object_version_ids}.
	// Experimental.
	SubObjectVersionIds *[]*float64 `field:"optional" json:"subObjectVersionIds" yaml:"subObjectVersionIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_version_id TfCatalogTable#view_version_id}.
	// Experimental.
	ViewVersionId *float64 `field:"optional" json:"viewVersionId" yaml:"viewVersionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_version_token TfCatalogTable#view_version_token}.
	// Experimental.
	ViewVersionToken *string `field:"optional" json:"viewVersionToken" yaml:"viewVersionToken"`
}

