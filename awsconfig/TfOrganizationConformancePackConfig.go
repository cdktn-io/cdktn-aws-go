package awsconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOrganizationConformancePackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#name TfOrganizationConformancePack#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#delivery_s3_bucket TfOrganizationConformancePack#delivery_s3_bucket}.
	// Experimental.
	DeliveryS3Bucket *string `field:"optional" json:"deliveryS3Bucket" yaml:"deliveryS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#delivery_s3_key_prefix TfOrganizationConformancePack#delivery_s3_key_prefix}.
	// Experimental.
	DeliveryS3KeyPrefix *string `field:"optional" json:"deliveryS3KeyPrefix" yaml:"deliveryS3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#excluded_accounts TfOrganizationConformancePack#excluded_accounts}.
	// Experimental.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#id TfOrganizationConformancePack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#input_parameter TfOrganizationConformancePack#input_parameter}
	// Experimental.
	InputParameter interface{} `field:"optional" json:"inputParameter" yaml:"inputParameter"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#region TfOrganizationConformancePack#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#template_body TfOrganizationConformancePack#template_body}.
	// Experimental.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#template_s3_uri TfOrganizationConformancePack#template_s3_uri}.
	// Experimental.
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#timeouts TfOrganizationConformancePack#timeouts}
	// Experimental.
	Timeouts *TfOrganizationConformancePack_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

