package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerTrainingJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#role_arn AwsSagemakerTrainingJob#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_job_name AwsSagemakerTrainingJob#training_job_name}.
	// Experimental.
	TrainingJobName *string `field:"required" json:"trainingJobName" yaml:"trainingJobName"`
	// algorithm_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#algorithm_specification AwsSagemakerTrainingJob#algorithm_specification}
	// Experimental.
	AlgorithmSpecification interface{} `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// checkpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#checkpoint_config AwsSagemakerTrainingJob#checkpoint_config}
	// Experimental.
	CheckpointConfig interface{} `field:"optional" json:"checkpointConfig" yaml:"checkpointConfig"`
	// debug_hook_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#debug_hook_config AwsSagemakerTrainingJob#debug_hook_config}
	// Experimental.
	DebugHookConfig interface{} `field:"optional" json:"debugHookConfig" yaml:"debugHookConfig"`
	// debug_rule_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#debug_rule_configurations AwsSagemakerTrainingJob#debug_rule_configurations}
	// Experimental.
	DebugRuleConfigurations interface{} `field:"optional" json:"debugRuleConfigurations" yaml:"debugRuleConfigurations"`
	// Whether to delete model packages in the configured model package group when destroying the training job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#delete_model_packages_on_destroy AwsSagemakerTrainingJob#delete_model_packages_on_destroy}
	// Experimental.
	DeleteModelPackagesOnDestroy interface{} `field:"optional" json:"deleteModelPackagesOnDestroy" yaml:"deleteModelPackagesOnDestroy"`
	// Whether to delete detached VPC ENIs that SageMaker may leave behind when destroying the training job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#delete_vpc_enis_on_destroy AwsSagemakerTrainingJob#delete_vpc_enis_on_destroy}
	// Experimental.
	DeleteVpcEnisOnDestroy interface{} `field:"optional" json:"deleteVpcEnisOnDestroy" yaml:"deleteVpcEnisOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#enable_inter_container_traffic_encryption AwsSagemakerTrainingJob#enable_inter_container_traffic_encryption}.
	// Experimental.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#enable_managed_spot_training AwsSagemakerTrainingJob#enable_managed_spot_training}.
	// Experimental.
	EnableManagedSpotTraining interface{} `field:"optional" json:"enableManagedSpotTraining" yaml:"enableManagedSpotTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#enable_network_isolation AwsSagemakerTrainingJob#enable_network_isolation}.
	// Experimental.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#environment AwsSagemakerTrainingJob#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// experiment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#experiment_config AwsSagemakerTrainingJob#experiment_config}
	// Experimental.
	ExperimentConfig interface{} `field:"optional" json:"experimentConfig" yaml:"experimentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#hyper_parameters AwsSagemakerTrainingJob#hyper_parameters}.
	// Experimental.
	HyperParameters *map[string]*string `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// infra_check_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#infra_check_config AwsSagemakerTrainingJob#infra_check_config}
	// Experimental.
	InfraCheckConfig interface{} `field:"optional" json:"infraCheckConfig" yaml:"infraCheckConfig"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#input_data_config AwsSagemakerTrainingJob#input_data_config}
	// Experimental.
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// mlflow_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_config AwsSagemakerTrainingJob#mlflow_config}
	// Experimental.
	MlflowConfig interface{} `field:"optional" json:"mlflowConfig" yaml:"mlflowConfig"`
	// model_package_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#model_package_config AwsSagemakerTrainingJob#model_package_config}
	// Experimental.
	ModelPackageConfig interface{} `field:"optional" json:"modelPackageConfig" yaml:"modelPackageConfig"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#output_data_config AwsSagemakerTrainingJob#output_data_config}
	// Experimental.
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// profiler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#profiler_config AwsSagemakerTrainingJob#profiler_config}
	// Experimental.
	ProfilerConfig interface{} `field:"optional" json:"profilerConfig" yaml:"profilerConfig"`
	// profiler_rule_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#profiler_rule_configurations AwsSagemakerTrainingJob#profiler_rule_configurations}
	// Experimental.
	ProfilerRuleConfigurations interface{} `field:"optional" json:"profilerRuleConfigurations" yaml:"profilerRuleConfigurations"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#region AwsSagemakerTrainingJob#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// remote_debug_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#remote_debug_config AwsSagemakerTrainingJob#remote_debug_config}
	// Experimental.
	RemoteDebugConfig interface{} `field:"optional" json:"remoteDebugConfig" yaml:"remoteDebugConfig"`
	// resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#resource_config AwsSagemakerTrainingJob#resource_config}
	// Experimental.
	ResourceConfig interface{} `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#retry_strategy AwsSagemakerTrainingJob#retry_strategy}
	// Experimental.
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// serverless_job_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#serverless_job_config AwsSagemakerTrainingJob#serverless_job_config}
	// Experimental.
	ServerlessJobConfig interface{} `field:"optional" json:"serverlessJobConfig" yaml:"serverlessJobConfig"`
	// session_chaining_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#session_chaining_config AwsSagemakerTrainingJob#session_chaining_config}
	// Experimental.
	SessionChainingConfig interface{} `field:"optional" json:"sessionChainingConfig" yaml:"sessionChainingConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#stopping_condition AwsSagemakerTrainingJob#stopping_condition}
	// Experimental.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#tags AwsSagemakerTrainingJob#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tensor_board_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#tensor_board_output_config AwsSagemakerTrainingJob#tensor_board_output_config}
	// Experimental.
	TensorBoardOutputConfig interface{} `field:"optional" json:"tensorBoardOutputConfig" yaml:"tensorBoardOutputConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#timeouts AwsSagemakerTrainingJob#timeouts}
	// Experimental.
	Timeouts *AwsSagemakerTrainingJob_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#vpc_config AwsSagemakerTrainingJob#vpc_config}
	// Experimental.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

