package awslakeformation


// Experimental.
type TfOptIn_LfTagPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#resource_type TfOptIn#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#catalog_id TfOptIn#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#expression TfOptIn#expression}.
	// Experimental.
	Expression *[]*string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#expression_name TfOptIn#expression_name}.
	// Experimental.
	ExpressionName *string `field:"optional" json:"expressionName" yaml:"expressionName"`
}

