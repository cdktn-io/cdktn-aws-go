package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNotebookInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#instance_type TfNotebookInstance#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#name TfNotebookInstance#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#role_arn TfNotebookInstance#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#additional_code_repositories TfNotebookInstance#additional_code_repositories}.
	// Experimental.
	AdditionalCodeRepositories *[]*string `field:"optional" json:"additionalCodeRepositories" yaml:"additionalCodeRepositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#default_code_repository TfNotebookInstance#default_code_repository}.
	// Experimental.
	DefaultCodeRepository *string `field:"optional" json:"defaultCodeRepository" yaml:"defaultCodeRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#direct_internet_access TfNotebookInstance#direct_internet_access}.
	// Experimental.
	DirectInternetAccess *string `field:"optional" json:"directInternetAccess" yaml:"directInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#id TfNotebookInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_metadata_service_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#instance_metadata_service_configuration TfNotebookInstance#instance_metadata_service_configuration}
	// Experimental.
	InstanceMetadataServiceConfiguration *TfNotebookInstance_InstanceMetadataServiceConfigurationProperty `field:"optional" json:"instanceMetadataServiceConfiguration" yaml:"instanceMetadataServiceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#kms_key_id TfNotebookInstance#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#lifecycle_config_name TfNotebookInstance#lifecycle_config_name}.
	// Experimental.
	LifecycleConfigName *string `field:"optional" json:"lifecycleConfigName" yaml:"lifecycleConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#platform_identifier TfNotebookInstance#platform_identifier}.
	// Experimental.
	PlatformIdentifier *string `field:"optional" json:"platformIdentifier" yaml:"platformIdentifier"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#region TfNotebookInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#root_access TfNotebookInstance#root_access}.
	// Experimental.
	RootAccess *string `field:"optional" json:"rootAccess" yaml:"rootAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#security_groups TfNotebookInstance#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#subnet_id TfNotebookInstance#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#tags TfNotebookInstance#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#tags_all TfNotebookInstance#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_notebook_instance#volume_size TfNotebookInstance#volume_size}.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

