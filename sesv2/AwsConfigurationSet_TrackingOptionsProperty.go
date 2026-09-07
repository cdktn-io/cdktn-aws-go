package sesv2


// Experimental.
type AwsConfigurationSet_TrackingOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#custom_redirect_domain AwsConfigurationSet#custom_redirect_domain}.
	// Experimental.
	CustomRedirectDomain *string `field:"required" json:"customRedirectDomain" yaml:"customRedirectDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#https_policy AwsConfigurationSet#https_policy}.
	// Experimental.
	HttpsPolicy *string `field:"optional" json:"httpsPolicy" yaml:"httpsPolicy"`
}

