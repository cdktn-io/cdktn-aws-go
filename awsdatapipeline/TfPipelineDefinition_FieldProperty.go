package awsdatapipeline


// Experimental.
type TfPipelineDefinition_FieldProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#key TfPipelineDefinition#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#ref_value TfPipelineDefinition#ref_value}.
	// Experimental.
	RefValue *string `field:"optional" json:"refValue" yaml:"refValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#string_value TfPipelineDefinition#string_value}.
	// Experimental.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

