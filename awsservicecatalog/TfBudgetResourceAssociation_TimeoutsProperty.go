package awsservicecatalog


// Experimental.
type TfBudgetResourceAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_budget_resource_association#create TfBudgetResourceAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_budget_resource_association#delete TfBudgetResourceAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_budget_resource_association#read TfBudgetResourceAssociation#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

