package sagemakerai


// Experimental.
type AwsUserProfile_StudioWebPortalSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#hidden_app_types AwsUserProfile#hidden_app_types}.
	// Experimental.
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#hidden_instance_types AwsUserProfile#hidden_instance_types}.
	// Experimental.
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#hidden_ml_tools AwsUserProfile#hidden_ml_tools}.
	// Experimental.
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
}

