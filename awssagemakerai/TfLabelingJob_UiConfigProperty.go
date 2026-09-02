package awssagemakerai


// Experimental.
type TfLabelingJob_UiConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#human_task_ui_arn TfLabelingJob#human_task_ui_arn}.
	// Experimental.
	HumanTaskUiArn *string `field:"optional" json:"humanTaskUiArn" yaml:"humanTaskUiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#ui_template_s3_uri TfLabelingJob#ui_template_s3_uri}.
	// Experimental.
	UiTemplateS3Uri *string `field:"optional" json:"uiTemplateS3Uri" yaml:"uiTemplateS3Uri"`
}

