package awslakeformation


// Experimental.
type TfOptIn_DataLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#resource_arn TfOptIn#resource_arn}.
	// Experimental.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#catalog_id TfOptIn#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

