package awsservicecatalog


// Experimental.
type TfServiceAction_DefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_service_action#name TfServiceAction#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_service_action#version TfServiceAction#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_service_action#assume_role TfServiceAction#assume_role}.
	// Experimental.
	AssumeRole *string `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_service_action#parameters TfServiceAction#parameters}.
	// Experimental.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_service_action#type TfServiceAction#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

