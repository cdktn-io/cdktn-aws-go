package lambda

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#function_name AwsFunction#function_name}.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#role AwsFunction#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#architectures AwsFunction#architectures}.
	// Experimental.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// capacity_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#capacity_provider_config AwsFunction#capacity_provider_config}
	// Experimental.
	CapacityProviderConfig *AwsFunction_CapacityProviderConfigProperty `field:"optional" json:"capacityProviderConfig" yaml:"capacityProviderConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#code_sha256 AwsFunction#code_sha256}.
	// Experimental.
	CodeSha256 *string `field:"optional" json:"codeSha256" yaml:"codeSha256"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#code_signing_config_arn AwsFunction#code_signing_config_arn}.
	// Experimental.
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#dead_letter_config AwsFunction#dead_letter_config}
	// Experimental.
	DeadLetterConfig *AwsFunction_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#description AwsFunction#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// durable_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#durable_config AwsFunction#durable_config}
	// Experimental.
	DurableConfig *AwsFunction_DurableConfigProperty `field:"optional" json:"durableConfig" yaml:"durableConfig"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#environment AwsFunction#environment}
	// Experimental.
	Environment *AwsFunction_EnvironmentProperty `field:"optional" json:"environment" yaml:"environment"`
	// ephemeral_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#ephemeral_storage AwsFunction#ephemeral_storage}
	// Experimental.
	EphemeralStorage *AwsFunction_EphemeralStorageProperty `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#filename AwsFunction#filename}.
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#file_system_config AwsFunction#file_system_config}
	// Experimental.
	FileSystemConfig *AwsFunction_FileSystemConfigProperty `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#handler AwsFunction#handler}.
	// Experimental.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#id AwsFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#image_config AwsFunction#image_config}
	// Experimental.
	ImageConfig *AwsFunction_ImageConfigProperty `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#image_uri AwsFunction#image_uri}.
	// Experimental.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#kms_key_arn AwsFunction#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#layers AwsFunction#layers}.
	// Experimental.
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#logging_config AwsFunction#logging_config}
	// Experimental.
	LoggingConfig *AwsFunction_LoggingConfigProperty `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#memory_size AwsFunction#memory_size}.
	// Experimental.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#package_type AwsFunction#package_type}.
	// Experimental.
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#publish AwsFunction#publish}.
	// Experimental.
	Publish interface{} `field:"optional" json:"publish" yaml:"publish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#publish_to AwsFunction#publish_to}.
	// Experimental.
	PublishTo *string `field:"optional" json:"publishTo" yaml:"publishTo"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#region AwsFunction#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#replacement_security_group_ids AwsFunction#replacement_security_group_ids}.
	// Experimental.
	ReplacementSecurityGroupIds *[]*string `field:"optional" json:"replacementSecurityGroupIds" yaml:"replacementSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#replace_security_groups_on_destroy AwsFunction#replace_security_groups_on_destroy}.
	// Experimental.
	ReplaceSecurityGroupsOnDestroy interface{} `field:"optional" json:"replaceSecurityGroupsOnDestroy" yaml:"replaceSecurityGroupsOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#reserved_concurrent_executions AwsFunction#reserved_concurrent_executions}.
	// Experimental.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#runtime AwsFunction#runtime}.
	// Experimental.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#s3_bucket AwsFunction#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#s3_key AwsFunction#s3_key}.
	// Experimental.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#s3_object_version AwsFunction#s3_object_version}.
	// Experimental.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#skip_destroy AwsFunction#skip_destroy}.
	// Experimental.
	SkipDestroy interface{} `field:"optional" json:"skipDestroy" yaml:"skipDestroy"`
	// snap_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#snap_start AwsFunction#snap_start}
	// Experimental.
	SnapStart *AwsFunction_SnapStartProperty `field:"optional" json:"snapStart" yaml:"snapStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#source_code_hash AwsFunction#source_code_hash}.
	// Experimental.
	SourceCodeHash *string `field:"optional" json:"sourceCodeHash" yaml:"sourceCodeHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#source_kms_key_arn AwsFunction#source_kms_key_arn}.
	// Experimental.
	SourceKmsKeyArn *string `field:"optional" json:"sourceKmsKeyArn" yaml:"sourceKmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#tags AwsFunction#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#tags_all AwsFunction#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tenancy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#tenancy_config AwsFunction#tenancy_config}
	// Experimental.
	TenancyConfig *AwsFunction_TenancyConfigProperty `field:"optional" json:"tenancyConfig" yaml:"tenancyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#timeout AwsFunction#timeout}.
	// Experimental.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#timeouts AwsFunction#timeouts}
	// Experimental.
	Timeouts *AwsFunction_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tracing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#tracing_config AwsFunction#tracing_config}
	// Experimental.
	TracingConfig *AwsFunction_TracingConfigProperty `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#use_resource_timeout_for_propagation AwsFunction#use_resource_timeout_for_propagation}.
	// Experimental.
	UseResourceTimeoutForPropagation interface{} `field:"optional" json:"useResourceTimeoutForPropagation" yaml:"useResourceTimeoutForPropagation"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#vpc_config AwsFunction#vpc_config}
	// Experimental.
	VpcConfig *AwsFunction_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

