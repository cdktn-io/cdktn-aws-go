package awslakeformation


// Experimental.
type TfOptIn_LfTagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#key TfOptIn#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#values TfOptIn#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#catalog_id TfOptIn#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

