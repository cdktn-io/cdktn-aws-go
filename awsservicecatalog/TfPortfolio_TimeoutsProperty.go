package awsservicecatalog


// Experimental.
type TfPortfolio_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_portfolio#create TfPortfolio#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_portfolio#delete TfPortfolio#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_portfolio#read TfPortfolio#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_portfolio#update TfPortfolio#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

