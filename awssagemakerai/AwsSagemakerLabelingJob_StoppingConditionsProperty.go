package awssagemakerai


// Experimental.
type AwsSagemakerLabelingJob_StoppingConditionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#max_human_labeled_object_count AwsSagemakerLabelingJob#max_human_labeled_object_count}.
	// Experimental.
	MaxHumanLabeledObjectCount *float64 `field:"optional" json:"maxHumanLabeledObjectCount" yaml:"maxHumanLabeledObjectCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#max_percentage_of_input_dataset_labeled AwsSagemakerLabelingJob#max_percentage_of_input_dataset_labeled}.
	// Experimental.
	MaxPercentageOfInputDatasetLabeled *float64 `field:"optional" json:"maxPercentageOfInputDatasetLabeled" yaml:"maxPercentageOfInputDatasetLabeled"`
}

