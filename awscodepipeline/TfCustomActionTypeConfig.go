package awscodepipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCustomActionTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#category TfCustomActionType#category}.
	// Experimental.
	Category *string `field:"required" json:"category" yaml:"category"`
	// input_artifact_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#input_artifact_details TfCustomActionType#input_artifact_details}
	// Experimental.
	InputArtifactDetails *TfCustomActionType_InputArtifactDetailsProperty `field:"required" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// output_artifact_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#output_artifact_details TfCustomActionType#output_artifact_details}
	// Experimental.
	OutputArtifactDetails *TfCustomActionType_OutputArtifactDetailsProperty `field:"required" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#provider_name TfCustomActionType#provider_name}.
	// Experimental.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#version TfCustomActionType#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// configuration_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#configuration_property TfCustomActionType#configuration_property}
	// Experimental.
	ConfigurationProperty interface{} `field:"optional" json:"configurationProperty" yaml:"configurationProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#id TfCustomActionType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#region TfCustomActionType#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#settings TfCustomActionType#settings}
	// Experimental.
	Settings *TfCustomActionType_SettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#tags TfCustomActionType#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#tags_all TfCustomActionType#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

