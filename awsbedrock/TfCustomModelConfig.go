package awsbedrock

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCustomModelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#base_model_identifier TfCustomModel#base_model_identifier}.
	// Experimental.
	BaseModelIdentifier *string `field:"required" json:"baseModelIdentifier" yaml:"baseModelIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#custom_model_name TfCustomModel#custom_model_name}.
	// Experimental.
	CustomModelName *string `field:"required" json:"customModelName" yaml:"customModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#hyperparameters TfCustomModel#hyperparameters}.
	// Experimental.
	Hyperparameters *map[string]*string `field:"required" json:"hyperparameters" yaml:"hyperparameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#job_name TfCustomModel#job_name}.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#role_arn TfCustomModel#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#customization_type TfCustomModel#customization_type}.
	// Experimental.
	CustomizationType *string `field:"optional" json:"customizationType" yaml:"customizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#custom_model_kms_key_id TfCustomModel#custom_model_kms_key_id}.
	// Experimental.
	CustomModelKmsKeyId *string `field:"optional" json:"customModelKmsKeyId" yaml:"customModelKmsKeyId"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#output_data_config TfCustomModel#output_data_config}
	// Experimental.
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#region TfCustomModel#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#tags TfCustomModel#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#timeouts TfCustomModel#timeouts}
	// Experimental.
	Timeouts *TfCustomModel_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// training_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#training_data_config TfCustomModel#training_data_config}
	// Experimental.
	TrainingDataConfig interface{} `field:"optional" json:"trainingDataConfig" yaml:"trainingDataConfig"`
	// validation_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#validation_data_config TfCustomModel#validation_data_config}
	// Experimental.
	ValidationDataConfig interface{} `field:"optional" json:"validationDataConfig" yaml:"validationDataConfig"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model#vpc_config TfCustomModel#vpc_config}
	// Experimental.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

