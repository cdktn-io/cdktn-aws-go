package awslambda

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLambdaEventSourceMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#function_name AwsLambdaEventSourceMapping#function_name}.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// amazon_managed_kafka_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#amazon_managed_kafka_event_source_config AwsLambdaEventSourceMapping#amazon_managed_kafka_event_source_config}
	// Experimental.
	AmazonManagedKafkaEventSourceConfig *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#batch_size AwsLambdaEventSourceMapping#batch_size}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#bisect_batch_on_function_error AwsLambdaEventSourceMapping#bisect_batch_on_function_error}.
	// Experimental.
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#destination_config AwsLambdaEventSourceMapping#destination_config}
	// Experimental.
	DestinationConfig *AwsLambdaEventSourceMapping_DestinationConfigProperty `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// document_db_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#document_db_event_source_config AwsLambdaEventSourceMapping#document_db_event_source_config}
	// Experimental.
	DocumentDbEventSourceConfig *AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty `field:"optional" json:"documentDbEventSourceConfig" yaml:"documentDbEventSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#enabled AwsLambdaEventSourceMapping#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#event_source_arn AwsLambdaEventSourceMapping#event_source_arn}.
	// Experimental.
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#filter_criteria AwsLambdaEventSourceMapping#filter_criteria}
	// Experimental.
	FilterCriteria *AwsLambdaEventSourceMapping_FilterCriteriaProperty `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#function_response_types AwsLambdaEventSourceMapping#function_response_types}.
	// Experimental.
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#id AwsLambdaEventSourceMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#kms_key_arn AwsLambdaEventSourceMapping#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_batching_window_in_seconds AwsLambdaEventSourceMapping#maximum_batching_window_in_seconds}.
	// Experimental.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_record_age_in_seconds AwsLambdaEventSourceMapping#maximum_record_age_in_seconds}.
	// Experimental.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_retry_attempts AwsLambdaEventSourceMapping#maximum_retry_attempts}.
	// Experimental.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// metrics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#metrics_config AwsLambdaEventSourceMapping#metrics_config}
	// Experimental.
	MetricsConfig *AwsLambdaEventSourceMapping_MetricsConfigProperty `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#parallelization_factor AwsLambdaEventSourceMapping#parallelization_factor}.
	// Experimental.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// provisioned_poller_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#provisioned_poller_config AwsLambdaEventSourceMapping#provisioned_poller_config}
	// Experimental.
	ProvisionedPollerConfig *AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty `field:"optional" json:"provisionedPollerConfig" yaml:"provisionedPollerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#queues AwsLambdaEventSourceMapping#queues}.
	// Experimental.
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#region AwsLambdaEventSourceMapping#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#scaling_config AwsLambdaEventSourceMapping#scaling_config}
	// Experimental.
	ScalingConfig *AwsLambdaEventSourceMapping_ScalingConfigProperty `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// self_managed_event_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#self_managed_event_source AwsLambdaEventSourceMapping#self_managed_event_source}
	// Experimental.
	SelfManagedEventSource *AwsLambdaEventSourceMapping_SelfManagedEventSourceProperty `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// self_managed_kafka_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#self_managed_kafka_event_source_config AwsLambdaEventSourceMapping#self_managed_kafka_event_source_config}
	// Experimental.
	SelfManagedKafkaEventSourceConfig *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// source_access_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#source_access_configuration AwsLambdaEventSourceMapping#source_access_configuration}
	// Experimental.
	SourceAccessConfiguration interface{} `field:"optional" json:"sourceAccessConfiguration" yaml:"sourceAccessConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#starting_position AwsLambdaEventSourceMapping#starting_position}.
	// Experimental.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#starting_position_timestamp AwsLambdaEventSourceMapping#starting_position_timestamp}.
	// Experimental.
	StartingPositionTimestamp *string `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#tags AwsLambdaEventSourceMapping#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#tags_all AwsLambdaEventSourceMapping#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#timeouts AwsLambdaEventSourceMapping#timeouts}
	// Experimental.
	Timeouts *AwsLambdaEventSourceMapping_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#topics AwsLambdaEventSourceMapping#topics}.
	// Experimental.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#tumbling_window_in_seconds AwsLambdaEventSourceMapping#tumbling_window_in_seconds}.
	// Experimental.
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#use_resource_timeout_for_propagation AwsLambdaEventSourceMapping#use_resource_timeout_for_propagation}.
	// Experimental.
	UseResourceTimeoutForPropagation interface{} `field:"optional" json:"useResourceTimeoutForPropagation" yaml:"useResourceTimeoutForPropagation"`
}

