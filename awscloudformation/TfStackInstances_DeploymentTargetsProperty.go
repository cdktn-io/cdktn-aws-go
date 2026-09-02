package awscloudformation


// Experimental.
type TfStackInstances_DeploymentTargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_instances#account_filter_type TfStackInstances#account_filter_type}.
	// Experimental.
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_instances#accounts TfStackInstances#accounts}.
	// Experimental.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_instances#accounts_url TfStackInstances#accounts_url}.
	// Experimental.
	AccountsUrl *string `field:"optional" json:"accountsUrl" yaml:"accountsUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_instances#organizational_unit_ids TfStackInstances#organizational_unit_ids}.
	// Experimental.
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

