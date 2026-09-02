package awsservicecatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProvisionedProductConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#name TfProvisionedProduct#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#accept_language TfProvisionedProduct#accept_language}.
	// Experimental.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#id TfProvisionedProduct#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#ignore_errors TfProvisionedProduct#ignore_errors}.
	// Experimental.
	IgnoreErrors interface{} `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#notification_arns TfProvisionedProduct#notification_arns}.
	// Experimental.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#path_id TfProvisionedProduct#path_id}.
	// Experimental.
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#path_name TfProvisionedProduct#path_name}.
	// Experimental.
	PathName *string `field:"optional" json:"pathName" yaml:"pathName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#product_id TfProvisionedProduct#product_id}.
	// Experimental.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#product_name TfProvisionedProduct#product_name}.
	// Experimental.
	ProductName *string `field:"optional" json:"productName" yaml:"productName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#provisioning_artifact_id TfProvisionedProduct#provisioning_artifact_id}.
	// Experimental.
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#provisioning_artifact_name TfProvisionedProduct#provisioning_artifact_name}.
	// Experimental.
	ProvisioningArtifactName *string `field:"optional" json:"provisioningArtifactName" yaml:"provisioningArtifactName"`
	// provisioning_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#provisioning_parameters TfProvisionedProduct#provisioning_parameters}
	// Experimental.
	ProvisioningParameters interface{} `field:"optional" json:"provisioningParameters" yaml:"provisioningParameters"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#region TfProvisionedProduct#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#retain_physical_resources TfProvisionedProduct#retain_physical_resources}.
	// Experimental.
	RetainPhysicalResources interface{} `field:"optional" json:"retainPhysicalResources" yaml:"retainPhysicalResources"`
	// stack_set_provisioning_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#stack_set_provisioning_preferences TfProvisionedProduct#stack_set_provisioning_preferences}
	// Experimental.
	StackSetProvisioningPreferences *TfProvisionedProduct_StackSetProvisioningPreferencesProperty `field:"optional" json:"stackSetProvisioningPreferences" yaml:"stackSetProvisioningPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#tags TfProvisionedProduct#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#tags_all TfProvisionedProduct#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#timeouts TfProvisionedProduct#timeouts}
	// Experimental.
	Timeouts *TfProvisionedProduct_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

