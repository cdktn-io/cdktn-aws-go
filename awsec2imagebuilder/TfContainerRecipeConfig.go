package awsec2imagebuilder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfContainerRecipeConfig struct {
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
	// component block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#component TfContainerRecipe#component}
	// Experimental.
	Component interface{} `field:"required" json:"component" yaml:"component"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#container_type TfContainerRecipe#container_type}.
	// Experimental.
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#name TfContainerRecipe#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#parent_image TfContainerRecipe#parent_image}.
	// Experimental.
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// target_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#target_repository TfContainerRecipe#target_repository}
	// Experimental.
	TargetRepository *TfContainerRecipe_TargetRepositoryProperty `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#version TfContainerRecipe#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#description TfContainerRecipe#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#dockerfile_template_data TfContainerRecipe#dockerfile_template_data}.
	// Experimental.
	DockerfileTemplateData *string `field:"optional" json:"dockerfileTemplateData" yaml:"dockerfileTemplateData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#dockerfile_template_uri TfContainerRecipe#dockerfile_template_uri}.
	// Experimental.
	DockerfileTemplateUri *string `field:"optional" json:"dockerfileTemplateUri" yaml:"dockerfileTemplateUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#id TfContainerRecipe#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#instance_configuration TfContainerRecipe#instance_configuration}
	// Experimental.
	InstanceConfiguration *TfContainerRecipe_InstanceConfigurationProperty `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#kms_key_id TfContainerRecipe#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#platform_override TfContainerRecipe#platform_override}.
	// Experimental.
	PlatformOverride *string `field:"optional" json:"platformOverride" yaml:"platformOverride"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#region TfContainerRecipe#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#tags TfContainerRecipe#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#tags_all TfContainerRecipe#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#working_directory TfContainerRecipe#working_directory}.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

