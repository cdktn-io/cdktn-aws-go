package awsglue


// Experimental.
type AwsGlueCatalogTable_ViewDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#definer AwsGlueCatalogTable#definer}.
	// Experimental.
	Definer *string `field:"optional" json:"definer" yaml:"definer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#is_protected AwsGlueCatalogTable#is_protected}.
	// Experimental.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#last_refresh_type AwsGlueCatalogTable#last_refresh_type}.
	// Experimental.
	LastRefreshType *string `field:"optional" json:"lastRefreshType" yaml:"lastRefreshType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#refresh_seconds AwsGlueCatalogTable#refresh_seconds}.
	// Experimental.
	RefreshSeconds *float64 `field:"optional" json:"refreshSeconds" yaml:"refreshSeconds"`
	// representations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#representations AwsGlueCatalogTable#representations}
	// Experimental.
	Representations interface{} `field:"optional" json:"representations" yaml:"representations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sub_objects AwsGlueCatalogTable#sub_objects}.
	// Experimental.
	SubObjects *[]*string `field:"optional" json:"subObjects" yaml:"subObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sub_object_version_ids AwsGlueCatalogTable#sub_object_version_ids}.
	// Experimental.
	SubObjectVersionIds *[]*float64 `field:"optional" json:"subObjectVersionIds" yaml:"subObjectVersionIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_version_id AwsGlueCatalogTable#view_version_id}.
	// Experimental.
	ViewVersionId *float64 `field:"optional" json:"viewVersionId" yaml:"viewVersionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_version_token AwsGlueCatalogTable#view_version_token}.
	// Experimental.
	ViewVersionToken *string `field:"optional" json:"viewVersionToken" yaml:"viewVersionToken"`
}

