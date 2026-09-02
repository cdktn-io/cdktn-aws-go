package awsinspector


// Experimental.
type TfOrganizationConfiguration_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#create TfOrganizationConfiguration#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#delete TfOrganizationConfiguration#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#update TfOrganizationConfiguration#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

