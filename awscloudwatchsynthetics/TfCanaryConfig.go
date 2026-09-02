package awscloudwatchsynthetics

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCanaryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#artifact_s3_location TfCanary#artifact_s3_location}.
	// Experimental.
	ArtifactS3Location *string `field:"required" json:"artifactS3Location" yaml:"artifactS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#execution_role_arn TfCanary#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#handler TfCanary#handler}.
	// Experimental.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#name TfCanary#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#runtime_version TfCanary#runtime_version}.
	// Experimental.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#schedule TfCanary#schedule}
	// Experimental.
	Schedule *TfCanary_ScheduleProperty `field:"required" json:"schedule" yaml:"schedule"`
	// artifact_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#artifact_config TfCanary#artifact_config}
	// Experimental.
	ArtifactConfig *TfCanary_ArtifactConfigProperty `field:"optional" json:"artifactConfig" yaml:"artifactConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#delete_lambda TfCanary#delete_lambda}.
	// Experimental.
	DeleteLambda interface{} `field:"optional" json:"deleteLambda" yaml:"deleteLambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#failure_retention_period TfCanary#failure_retention_period}.
	// Experimental.
	FailureRetentionPeriod *float64 `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#id TfCanary#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#region TfCanary#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// run_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#run_config TfCanary#run_config}
	// Experimental.
	RunConfig *TfCanary_RunConfigProperty `field:"optional" json:"runConfig" yaml:"runConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#s3_bucket TfCanary#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#s3_key TfCanary#s3_key}.
	// Experimental.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#s3_version TfCanary#s3_version}.
	// Experimental.
	S3Version *string `field:"optional" json:"s3Version" yaml:"s3Version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#start_canary TfCanary#start_canary}.
	// Experimental.
	StartCanary interface{} `field:"optional" json:"startCanary" yaml:"startCanary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#success_retention_period TfCanary#success_retention_period}.
	// Experimental.
	SuccessRetentionPeriod *float64 `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#tags TfCanary#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#tags_all TfCanary#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#vpc_config TfCanary#vpc_config}
	// Experimental.
	VpcConfig *TfCanary_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#zip_file TfCanary#zip_file}.
	// Experimental.
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

