package awssqs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#content_based_deduplication TfQueue#content_based_deduplication}.
	// Experimental.
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#deduplication_scope TfQueue#deduplication_scope}.
	// Experimental.
	DeduplicationScope *string `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#delay_seconds TfQueue#delay_seconds}.
	// Experimental.
	DelaySeconds *float64 `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#fifo_queue TfQueue#fifo_queue}.
	// Experimental.
	FifoQueue interface{} `field:"optional" json:"fifoQueue" yaml:"fifoQueue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#fifo_throughput_limit TfQueue#fifo_throughput_limit}.
	// Experimental.
	FifoThroughputLimit *string `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#id TfQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#kms_data_key_reuse_period_seconds TfQueue#kms_data_key_reuse_period_seconds}.
	// Experimental.
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#kms_master_key_id TfQueue#kms_master_key_id}.
	// Experimental.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#max_message_size TfQueue#max_message_size}.
	// Experimental.
	MaxMessageSize *float64 `field:"optional" json:"maxMessageSize" yaml:"maxMessageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#message_retention_seconds TfQueue#message_retention_seconds}.
	// Experimental.
	MessageRetentionSeconds *float64 `field:"optional" json:"messageRetentionSeconds" yaml:"messageRetentionSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#name TfQueue#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#name_prefix TfQueue#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#policy TfQueue#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#receive_wait_time_seconds TfQueue#receive_wait_time_seconds}.
	// Experimental.
	ReceiveWaitTimeSeconds *float64 `field:"optional" json:"receiveWaitTimeSeconds" yaml:"receiveWaitTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#redrive_allow_policy TfQueue#redrive_allow_policy}.
	// Experimental.
	RedriveAllowPolicy *string `field:"optional" json:"redriveAllowPolicy" yaml:"redriveAllowPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#redrive_policy TfQueue#redrive_policy}.
	// Experimental.
	RedrivePolicy *string `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#region TfQueue#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#sqs_managed_sse_enabled TfQueue#sqs_managed_sse_enabled}.
	// Experimental.
	SqsManagedSseEnabled interface{} `field:"optional" json:"sqsManagedSseEnabled" yaml:"sqsManagedSseEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#tags TfQueue#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#tags_all TfQueue#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#timeouts TfQueue#timeouts}
	// Experimental.
	Timeouts *TfQueue_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue#visibility_timeout_seconds TfQueue#visibility_timeout_seconds}.
	// Experimental.
	VisibilityTimeoutSeconds *float64 `field:"optional" json:"visibilityTimeoutSeconds" yaml:"visibilityTimeoutSeconds"`
}

