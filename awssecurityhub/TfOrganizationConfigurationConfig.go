package awssecurityhub

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOrganizationConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_organization_configuration#auto_enable TfOrganizationConfiguration#auto_enable}.
	// Experimental.
	AutoEnable interface{} `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_organization_configuration#auto_enable_standards TfOrganizationConfiguration#auto_enable_standards}.
	// Experimental.
	AutoEnableStandards *string `field:"optional" json:"autoEnableStandards" yaml:"autoEnableStandards"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_organization_configuration#id TfOrganizationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// organization_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_organization_configuration#organization_configuration TfOrganizationConfiguration#organization_configuration}
	// Experimental.
	OrganizationConfiguration *TfOrganizationConfiguration_OrganizationConfigurationProperty `field:"optional" json:"organizationConfiguration" yaml:"organizationConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_organization_configuration#region TfOrganizationConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_organization_configuration#timeouts TfOrganizationConfiguration#timeouts}
	// Experimental.
	Timeouts *TfOrganizationConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

