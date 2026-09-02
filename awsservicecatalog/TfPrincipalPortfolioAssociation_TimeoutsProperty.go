package awsservicecatalog


// Experimental.
type TfPrincipalPortfolioAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_principal_portfolio_association#create TfPrincipalPortfolioAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_principal_portfolio_association#delete TfPrincipalPortfolioAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_principal_portfolio_association#read TfPrincipalPortfolioAssociation#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

