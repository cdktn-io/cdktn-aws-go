package eventbridge

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#arn AwsTarget#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#rule AwsTarget#rule}.
	// Experimental.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// appsync_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#appsync_target AwsTarget#appsync_target}
	// Experimental.
	AppsyncTarget *AwsTarget_AppsyncTargetProperty `field:"optional" json:"appsyncTarget" yaml:"appsyncTarget"`
	// batch_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#batch_target AwsTarget#batch_target}
	// Experimental.
	BatchTarget *AwsTarget_BatchTargetProperty `field:"optional" json:"batchTarget" yaml:"batchTarget"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#dead_letter_config AwsTarget#dead_letter_config}
	// Experimental.
	DeadLetterConfig *AwsTarget_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// ecs_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#ecs_target AwsTarget#ecs_target}
	// Experimental.
	EcsTarget *AwsTarget_EcsTargetProperty `field:"optional" json:"ecsTarget" yaml:"ecsTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#event_bus_name AwsTarget#event_bus_name}.
	// Experimental.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#force_destroy AwsTarget#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// http_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#http_target AwsTarget#http_target}
	// Experimental.
	HttpTarget *AwsTarget_HttpTargetProperty `field:"optional" json:"httpTarget" yaml:"httpTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#id AwsTarget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#input AwsTarget#input}.
	// Experimental.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#input_path AwsTarget#input_path}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// input_transformer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#input_transformer AwsTarget#input_transformer}
	// Experimental.
	InputTransformer *AwsTarget_InputTransformerProperty `field:"optional" json:"inputTransformer" yaml:"inputTransformer"`
	// kinesis_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#kinesis_target AwsTarget#kinesis_target}
	// Experimental.
	KinesisTarget *AwsTarget_KinesisTargetProperty `field:"optional" json:"kinesisTarget" yaml:"kinesisTarget"`
	// redshift_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#redshift_target AwsTarget#redshift_target}
	// Experimental.
	RedshiftTarget *AwsTarget_RedshiftTargetProperty `field:"optional" json:"redshiftTarget" yaml:"redshiftTarget"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#region AwsTarget#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#retry_policy AwsTarget#retry_policy}
	// Experimental.
	RetryPolicy *AwsTarget_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#role_arn AwsTarget#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// run_command_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#run_command_targets AwsTarget#run_command_targets}
	// Experimental.
	RunCommandTargets interface{} `field:"optional" json:"runCommandTargets" yaml:"runCommandTargets"`
	// sagemaker_pipeline_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#sagemaker_pipeline_target AwsTarget#sagemaker_pipeline_target}
	// Experimental.
	SagemakerPipelineTarget *AwsTarget_SagemakerPipelineTargetProperty `field:"optional" json:"sagemakerPipelineTarget" yaml:"sagemakerPipelineTarget"`
	// sqs_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#sqs_target AwsTarget#sqs_target}
	// Experimental.
	SqsTarget *AwsTarget_SqsTargetProperty `field:"optional" json:"sqsTarget" yaml:"sqsTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#target_id AwsTarget#target_id}.
	// Experimental.
	TargetId *string `field:"optional" json:"targetId" yaml:"targetId"`
}

