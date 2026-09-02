package awssesv2


// Experimental.
type TfConfigurationSet_TrackingOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#custom_redirect_domain TfConfigurationSet#custom_redirect_domain}.
	// Experimental.
	CustomRedirectDomain *string `field:"required" json:"customRedirectDomain" yaml:"customRedirectDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#https_policy TfConfigurationSet#https_policy}.
	// Experimental.
	HttpsPolicy *string `field:"optional" json:"httpsPolicy" yaml:"httpsPolicy"`
}

