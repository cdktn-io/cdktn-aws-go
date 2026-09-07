package sagemakerai


// Experimental.
type AwsAlgorithm_TrainingSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_training_instance_types AwsAlgorithm#supported_training_instance_types}.
	// Experimental.
	SupportedTrainingInstanceTypes *[]*string `field:"required" json:"supportedTrainingInstanceTypes" yaml:"supportedTrainingInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#training_image AwsAlgorithm#training_image}.
	// Experimental.
	TrainingImage *string `field:"required" json:"trainingImage" yaml:"trainingImage"`
	// additional_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#additional_s3_data_source AwsAlgorithm#additional_s3_data_source}
	// Experimental.
	AdditionalS3DataSource interface{} `field:"optional" json:"additionalS3DataSource" yaml:"additionalS3DataSource"`
	// metric_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#metric_definitions AwsAlgorithm#metric_definitions}
	// Experimental.
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// supported_hyper_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_hyper_parameters AwsAlgorithm#supported_hyper_parameters}
	// Experimental.
	SupportedHyperParameters interface{} `field:"optional" json:"supportedHyperParameters" yaml:"supportedHyperParameters"`
	// supported_tuning_job_objective_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_tuning_job_objective_metrics AwsAlgorithm#supported_tuning_job_objective_metrics}
	// Experimental.
	SupportedTuningJobObjectiveMetrics interface{} `field:"optional" json:"supportedTuningJobObjectiveMetrics" yaml:"supportedTuningJobObjectiveMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supports_distributed_training AwsAlgorithm#supports_distributed_training}.
	// Experimental.
	SupportsDistributedTraining interface{} `field:"optional" json:"supportsDistributedTraining" yaml:"supportsDistributedTraining"`
	// training_channels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#training_channels AwsAlgorithm#training_channels}
	// Experimental.
	TrainingChannels interface{} `field:"optional" json:"trainingChannels" yaml:"trainingChannels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#training_image_digest AwsAlgorithm#training_image_digest}.
	// Experimental.
	TrainingImageDigest *string `field:"optional" json:"trainingImageDigest" yaml:"trainingImageDigest"`
}

