package awscloudformation


// Experimental.
type TfStackSetInstance_DeploymentTargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#account_filter_type TfStackSetInstance#account_filter_type}.
	// Experimental.
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#accounts TfStackSetInstance#accounts}.
	// Experimental.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#accounts_url TfStackSetInstance#accounts_url}.
	// Experimental.
	AccountsUrl *string `field:"optional" json:"accountsUrl" yaml:"accountsUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#organizational_unit_ids TfStackSetInstance#organizational_unit_ids}.
	// Experimental.
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

