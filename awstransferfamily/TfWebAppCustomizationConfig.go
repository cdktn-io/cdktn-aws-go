package awstransferfamily

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAppCustomizationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_web_app_customization#web_app_id TfWebAppCustomization#web_app_id}.
	// Experimental.
	WebAppId *string `field:"required" json:"webAppId" yaml:"webAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_web_app_customization#favicon_file TfWebAppCustomization#favicon_file}.
	// Experimental.
	FaviconFile *string `field:"optional" json:"faviconFile" yaml:"faviconFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_web_app_customization#logo_file TfWebAppCustomization#logo_file}.
	// Experimental.
	LogoFile *string `field:"optional" json:"logoFile" yaml:"logoFile"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_web_app_customization#region TfWebAppCustomization#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_web_app_customization#title TfWebAppCustomization#title}.
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

