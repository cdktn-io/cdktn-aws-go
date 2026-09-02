package awssagemakerai


// Experimental.
type TfUserProfile_StudioWebPortalSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#hidden_app_types TfUserProfile#hidden_app_types}.
	// Experimental.
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#hidden_instance_types TfUserProfile#hidden_instance_types}.
	// Experimental.
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#hidden_ml_tools TfUserProfile#hidden_ml_tools}.
	// Experimental.
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
}

