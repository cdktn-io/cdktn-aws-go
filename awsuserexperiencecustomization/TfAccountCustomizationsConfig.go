package awsuserexperiencecustomization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAccountCustomizationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/uxc_account_customizations#account_color TfAccountCustomizations#account_color}.
	// Experimental.
	AccountColor *string `field:"optional" json:"accountColor" yaml:"accountColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/uxc_account_customizations#visible_regions TfAccountCustomizations#visible_regions}.
	// Experimental.
	VisibleRegions *[]*string `field:"optional" json:"visibleRegions" yaml:"visibleRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/uxc_account_customizations#visible_services TfAccountCustomizations#visible_services}.
	// Experimental.
	VisibleServices *[]*string `field:"optional" json:"visibleServices" yaml:"visibleServices"`
}

