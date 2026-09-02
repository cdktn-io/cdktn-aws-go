package awscodepipeline


// Experimental.
type TfCodepipeline_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#category TfCodepipeline#category}.
	// Experimental.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#name TfCodepipeline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#owner TfCodepipeline#owner}.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#provider TfCodepipeline#provider}.
	// Experimental.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#version TfCodepipeline#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#commands TfCodepipeline#commands}.
	// Experimental.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#configuration TfCodepipeline#configuration}.
	// Experimental.
	Configuration *map[string]*string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#input_artifacts TfCodepipeline#input_artifacts}.
	// Experimental.
	InputArtifacts *[]*string `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#namespace TfCodepipeline#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#output_artifacts TfCodepipeline#output_artifacts}.
	// Experimental.
	OutputArtifacts *[]*string `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// output_artifacts_for_compute_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#output_artifacts_for_compute_action TfCodepipeline#output_artifacts_for_compute_action}
	// Experimental.
	OutputArtifactsForComputeAction interface{} `field:"optional" json:"outputArtifactsForComputeAction" yaml:"outputArtifactsForComputeAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#output_variables TfCodepipeline#output_variables}.
	// Experimental.
	OutputVariables *[]*string `field:"optional" json:"outputVariables" yaml:"outputVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#region TfCodepipeline#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#role_arn TfCodepipeline#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#run_order TfCodepipeline#run_order}.
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#timeout_in_minutes TfCodepipeline#timeout_in_minutes}.
	// Experimental.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

