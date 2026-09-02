package awscomprehend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEntityRecognizerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#data_access_role_arn TfEntityRecognizer#data_access_role_arn}.
	// Experimental.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#input_data_config TfEntityRecognizer#input_data_config}
	// Experimental.
	InputDataConfig *TfEntityRecognizer_InputDataConfigProperty `field:"required" json:"inputDataConfig" yaml:"inputDataConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#language_code TfEntityRecognizer#language_code}.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#name TfEntityRecognizer#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#id TfEntityRecognizer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#model_kms_key_id TfEntityRecognizer#model_kms_key_id}.
	// Experimental.
	ModelKmsKeyId *string `field:"optional" json:"modelKmsKeyId" yaml:"modelKmsKeyId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#region TfEntityRecognizer#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#tags TfEntityRecognizer#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#tags_all TfEntityRecognizer#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#timeouts TfEntityRecognizer#timeouts}
	// Experimental.
	Timeouts *TfEntityRecognizer_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#version_name TfEntityRecognizer#version_name}.
	// Experimental.
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#version_name_prefix TfEntityRecognizer#version_name_prefix}.
	// Experimental.
	VersionNamePrefix *string `field:"optional" json:"versionNamePrefix" yaml:"versionNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#volume_kms_key_id TfEntityRecognizer#volume_kms_key_id}.
	// Experimental.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#vpc_config TfEntityRecognizer#vpc_config}
	// Experimental.
	VpcConfig *TfEntityRecognizer_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

