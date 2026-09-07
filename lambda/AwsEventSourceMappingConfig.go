package lambda

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEventSourceMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#function_name AwsEventSourceMapping#function_name}.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// amazon_managed_kafka_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#amazon_managed_kafka_event_source_config AwsEventSourceMapping#amazon_managed_kafka_event_source_config}
	// Experimental.
	AmazonManagedKafkaEventSourceConfig *AwsEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#batch_size AwsEventSourceMapping#batch_size}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#bisect_batch_on_function_error AwsEventSourceMapping#bisect_batch_on_function_error}.
	// Experimental.
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#destination_config AwsEventSourceMapping#destination_config}
	// Experimental.
	DestinationConfig *AwsEventSourceMapping_DestinationConfigProperty `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// document_db_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#document_db_event_source_config AwsEventSourceMapping#document_db_event_source_config}
	// Experimental.
	DocumentDbEventSourceConfig *AwsEventSourceMapping_DocumentDbEventSourceConfigProperty `field:"optional" json:"documentDbEventSourceConfig" yaml:"documentDbEventSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#enabled AwsEventSourceMapping#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#event_source_arn AwsEventSourceMapping#event_source_arn}.
	// Experimental.
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#filter_criteria AwsEventSourceMapping#filter_criteria}
	// Experimental.
	FilterCriteria *AwsEventSourceMapping_FilterCriteriaProperty `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#function_response_types AwsEventSourceMapping#function_response_types}.
	// Experimental.
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#id AwsEventSourceMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#kms_key_arn AwsEventSourceMapping#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_batching_window_in_seconds AwsEventSourceMapping#maximum_batching_window_in_seconds}.
	// Experimental.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_record_age_in_seconds AwsEventSourceMapping#maximum_record_age_in_seconds}.
	// Experimental.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#maximum_retry_attempts AwsEventSourceMapping#maximum_retry_attempts}.
	// Experimental.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// metrics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#metrics_config AwsEventSourceMapping#metrics_config}
	// Experimental.
	MetricsConfig *AwsEventSourceMapping_MetricsConfigProperty `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#parallelization_factor AwsEventSourceMapping#parallelization_factor}.
	// Experimental.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// provisioned_poller_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#provisioned_poller_config AwsEventSourceMapping#provisioned_poller_config}
	// Experimental.
	ProvisionedPollerConfig *AwsEventSourceMapping_ProvisionedPollerConfigProperty `field:"optional" json:"provisionedPollerConfig" yaml:"provisionedPollerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#queues AwsEventSourceMapping#queues}.
	// Experimental.
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#region AwsEventSourceMapping#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#scaling_config AwsEventSourceMapping#scaling_config}
	// Experimental.
	ScalingConfig *AwsEventSourceMapping_ScalingConfigProperty `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// self_managed_event_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#self_managed_event_source AwsEventSourceMapping#self_managed_event_source}
	// Experimental.
	SelfManagedEventSource *AwsEventSourceMapping_SelfManagedEventSourceProperty `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// self_managed_kafka_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#self_managed_kafka_event_source_config AwsEventSourceMapping#self_managed_kafka_event_source_config}
	// Experimental.
	SelfManagedKafkaEventSourceConfig *AwsEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// source_access_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#source_access_configuration AwsEventSourceMapping#source_access_configuration}
	// Experimental.
	SourceAccessConfiguration interface{} `field:"optional" json:"sourceAccessConfiguration" yaml:"sourceAccessConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#starting_position AwsEventSourceMapping#starting_position}.
	// Experimental.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#starting_position_timestamp AwsEventSourceMapping#starting_position_timestamp}.
	// Experimental.
	StartingPositionTimestamp *string `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#tags AwsEventSourceMapping#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#tags_all AwsEventSourceMapping#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#timeouts AwsEventSourceMapping#timeouts}
	// Experimental.
	Timeouts *AwsEventSourceMapping_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#topics AwsEventSourceMapping#topics}.
	// Experimental.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#tumbling_window_in_seconds AwsEventSourceMapping#tumbling_window_in_seconds}.
	// Experimental.
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#use_resource_timeout_for_propagation AwsEventSourceMapping#use_resource_timeout_for_propagation}.
	// Experimental.
	UseResourceTimeoutForPropagation interface{} `field:"optional" json:"useResourceTimeoutForPropagation" yaml:"useResourceTimeoutForPropagation"`
}

