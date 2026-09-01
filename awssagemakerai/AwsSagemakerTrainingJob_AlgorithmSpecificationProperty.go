package awssagemakerai


// Experimental.
type AwsSagemakerTrainingJob_AlgorithmSpecificationProperty struct {
	// Name or ARN of a SageMaker algorithm resource. Exactly one of `algorithm_name` or `training_image` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#algorithm_name AwsSagemakerTrainingJob#algorithm_name}
	// Experimental.
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#container_arguments AwsSagemakerTrainingJob#container_arguments}.
	// Experimental.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#container_entrypoint AwsSagemakerTrainingJob#container_entrypoint}.
	// Experimental.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Whether SageMaker AI should publish time-series metrics.
	//
	// SageMaker enables this automatically for built-in algorithms, supported prebuilt images, and jobs with explicit `metric_definitions`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#enable_sagemaker_metrics_time_series AwsSagemakerTrainingJob#enable_sagemaker_metrics_time_series}
	// Experimental.
	EnableSagemakerMetricsTimeSeries interface{} `field:"optional" json:"enableSagemakerMetricsTimeSeries" yaml:"enableSagemakerMetricsTimeSeries"`
	// metric_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#metric_definitions AwsSagemakerTrainingJob#metric_definitions}
	// Experimental.
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Registry path of the training image.
	//
	// Exactly one of `algorithm_name` or `training_image` must be set. Use `metric_definitions` only when you need to extract custom metrics from your own training container logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_image AwsSagemakerTrainingJob#training_image}
	// Experimental.
	TrainingImage *string `field:"optional" json:"trainingImage" yaml:"trainingImage"`
	// training_image_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_image_config AwsSagemakerTrainingJob#training_image_config}
	// Experimental.
	TrainingImageConfig interface{} `field:"optional" json:"trainingImageConfig" yaml:"trainingImageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_input_mode AwsSagemakerTrainingJob#training_input_mode}.
	// Experimental.
	TrainingInputMode *string `field:"optional" json:"trainingInputMode" yaml:"trainingInputMode"`
}

