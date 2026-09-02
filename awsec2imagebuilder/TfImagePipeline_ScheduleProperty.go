package awsec2imagebuilder


// Experimental.
type TfImagePipeline_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#schedule_expression TfImagePipeline#schedule_expression}.
	// Experimental.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#pipeline_execution_start_condition TfImagePipeline#pipeline_execution_start_condition}.
	// Experimental.
	PipelineExecutionStartCondition *string `field:"optional" json:"pipelineExecutionStartCondition" yaml:"pipelineExecutionStartCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#timezone TfImagePipeline#timezone}.
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

