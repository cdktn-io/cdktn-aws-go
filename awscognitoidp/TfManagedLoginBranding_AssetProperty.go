package awscognitoidp


// Experimental.
type TfManagedLoginBranding_AssetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_login_branding#category TfManagedLoginBranding#category}.
	// Experimental.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_login_branding#color_mode TfManagedLoginBranding#color_mode}.
	// Experimental.
	ColorMode *string `field:"required" json:"colorMode" yaml:"colorMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_login_branding#extension TfManagedLoginBranding#extension}.
	// Experimental.
	Extension *string `field:"required" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_login_branding#bytes TfManagedLoginBranding#bytes}.
	// Experimental.
	Bytes *string `field:"optional" json:"bytes" yaml:"bytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_login_branding#resource_id TfManagedLoginBranding#resource_id}.
	// Experimental.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

