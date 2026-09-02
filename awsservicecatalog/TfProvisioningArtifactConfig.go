package awsservicecatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProvisioningArtifactConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#product_id TfProvisioningArtifact#product_id}.
	// Experimental.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#accept_language TfProvisioningArtifact#accept_language}.
	// Experimental.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#active TfProvisioningArtifact#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#description TfProvisioningArtifact#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#disable_template_validation TfProvisioningArtifact#disable_template_validation}.
	// Experimental.
	DisableTemplateValidation interface{} `field:"optional" json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#guidance TfProvisioningArtifact#guidance}.
	// Experimental.
	Guidance *string `field:"optional" json:"guidance" yaml:"guidance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#id TfProvisioningArtifact#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#name TfProvisioningArtifact#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#region TfProvisioningArtifact#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#template_physical_id TfProvisioningArtifact#template_physical_id}.
	// Experimental.
	TemplatePhysicalId *string `field:"optional" json:"templatePhysicalId" yaml:"templatePhysicalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#template_url TfProvisioningArtifact#template_url}.
	// Experimental.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#timeouts TfProvisioningArtifact#timeouts}
	// Experimental.
	Timeouts *TfProvisioningArtifact_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioning_artifact#type TfProvisioningArtifact#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

