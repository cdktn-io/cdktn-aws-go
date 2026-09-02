package awssns

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTopicConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#application_failure_feedback_role_arn TfTopic#application_failure_feedback_role_arn}.
	// Experimental.
	ApplicationFailureFeedbackRoleArn *string `field:"optional" json:"applicationFailureFeedbackRoleArn" yaml:"applicationFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#application_success_feedback_role_arn TfTopic#application_success_feedback_role_arn}.
	// Experimental.
	ApplicationSuccessFeedbackRoleArn *string `field:"optional" json:"applicationSuccessFeedbackRoleArn" yaml:"applicationSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#application_success_feedback_sample_rate TfTopic#application_success_feedback_sample_rate}.
	// Experimental.
	ApplicationSuccessFeedbackSampleRate *float64 `field:"optional" json:"applicationSuccessFeedbackSampleRate" yaml:"applicationSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#archive_policy TfTopic#archive_policy}.
	// Experimental.
	ArchivePolicy *string `field:"optional" json:"archivePolicy" yaml:"archivePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#content_based_deduplication TfTopic#content_based_deduplication}.
	// Experimental.
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#delivery_policy TfTopic#delivery_policy}.
	// Experimental.
	DeliveryPolicy *string `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#display_name TfTopic#display_name}.
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#fifo_throughput_scope TfTopic#fifo_throughput_scope}.
	// Experimental.
	FifoThroughputScope *string `field:"optional" json:"fifoThroughputScope" yaml:"fifoThroughputScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#fifo_topic TfTopic#fifo_topic}.
	// Experimental.
	FifoTopic interface{} `field:"optional" json:"fifoTopic" yaml:"fifoTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#firehose_failure_feedback_role_arn TfTopic#firehose_failure_feedback_role_arn}.
	// Experimental.
	FirehoseFailureFeedbackRoleArn *string `field:"optional" json:"firehoseFailureFeedbackRoleArn" yaml:"firehoseFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#firehose_success_feedback_role_arn TfTopic#firehose_success_feedback_role_arn}.
	// Experimental.
	FirehoseSuccessFeedbackRoleArn *string `field:"optional" json:"firehoseSuccessFeedbackRoleArn" yaml:"firehoseSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#firehose_success_feedback_sample_rate TfTopic#firehose_success_feedback_sample_rate}.
	// Experimental.
	FirehoseSuccessFeedbackSampleRate *float64 `field:"optional" json:"firehoseSuccessFeedbackSampleRate" yaml:"firehoseSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#http_failure_feedback_role_arn TfTopic#http_failure_feedback_role_arn}.
	// Experimental.
	HttpFailureFeedbackRoleArn *string `field:"optional" json:"httpFailureFeedbackRoleArn" yaml:"httpFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#http_success_feedback_role_arn TfTopic#http_success_feedback_role_arn}.
	// Experimental.
	HttpSuccessFeedbackRoleArn *string `field:"optional" json:"httpSuccessFeedbackRoleArn" yaml:"httpSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#http_success_feedback_sample_rate TfTopic#http_success_feedback_sample_rate}.
	// Experimental.
	HttpSuccessFeedbackSampleRate *float64 `field:"optional" json:"httpSuccessFeedbackSampleRate" yaml:"httpSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#id TfTopic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#kms_master_key_id TfTopic#kms_master_key_id}.
	// Experimental.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#lambda_failure_feedback_role_arn TfTopic#lambda_failure_feedback_role_arn}.
	// Experimental.
	LambdaFailureFeedbackRoleArn *string `field:"optional" json:"lambdaFailureFeedbackRoleArn" yaml:"lambdaFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#lambda_success_feedback_role_arn TfTopic#lambda_success_feedback_role_arn}.
	// Experimental.
	LambdaSuccessFeedbackRoleArn *string `field:"optional" json:"lambdaSuccessFeedbackRoleArn" yaml:"lambdaSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#lambda_success_feedback_sample_rate TfTopic#lambda_success_feedback_sample_rate}.
	// Experimental.
	LambdaSuccessFeedbackSampleRate *float64 `field:"optional" json:"lambdaSuccessFeedbackSampleRate" yaml:"lambdaSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#name TfTopic#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#name_prefix TfTopic#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#policy TfTopic#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#region TfTopic#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#signature_version TfTopic#signature_version}.
	// Experimental.
	SignatureVersion *float64 `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#sqs_failure_feedback_role_arn TfTopic#sqs_failure_feedback_role_arn}.
	// Experimental.
	SqsFailureFeedbackRoleArn *string `field:"optional" json:"sqsFailureFeedbackRoleArn" yaml:"sqsFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#sqs_success_feedback_role_arn TfTopic#sqs_success_feedback_role_arn}.
	// Experimental.
	SqsSuccessFeedbackRoleArn *string `field:"optional" json:"sqsSuccessFeedbackRoleArn" yaml:"sqsSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#sqs_success_feedback_sample_rate TfTopic#sqs_success_feedback_sample_rate}.
	// Experimental.
	SqsSuccessFeedbackSampleRate *float64 `field:"optional" json:"sqsSuccessFeedbackSampleRate" yaml:"sqsSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#tags TfTopic#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#tags_all TfTopic#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic#tracing_config TfTopic#tracing_config}.
	// Experimental.
	TracingConfig *string `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
}

