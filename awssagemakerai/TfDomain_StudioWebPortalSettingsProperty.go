package awssagemakerai


// Experimental.
type TfDomain_StudioWebPortalSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#hidden_app_types TfDomain#hidden_app_types}.
	// Experimental.
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#hidden_instance_types TfDomain#hidden_instance_types}.
	// Experimental.
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#hidden_ml_tools TfDomain#hidden_ml_tools}.
	// Experimental.
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
}

