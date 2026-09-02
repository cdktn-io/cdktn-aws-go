package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStudioLifecycleConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#studio_lifecycle_config_app_type TfStudioLifecycleConfig#studio_lifecycle_config_app_type}.
	// Experimental.
	StudioLifecycleConfigAppType *string `field:"required" json:"studioLifecycleConfigAppType" yaml:"studioLifecycleConfigAppType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#studio_lifecycle_config_content TfStudioLifecycleConfig#studio_lifecycle_config_content}.
	// Experimental.
	StudioLifecycleConfigContent *string `field:"required" json:"studioLifecycleConfigContent" yaml:"studioLifecycleConfigContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#studio_lifecycle_config_name TfStudioLifecycleConfig#studio_lifecycle_config_name}.
	// Experimental.
	StudioLifecycleConfigName *string `field:"required" json:"studioLifecycleConfigName" yaml:"studioLifecycleConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#id TfStudioLifecycleConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#region TfStudioLifecycleConfig#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#tags TfStudioLifecycleConfig#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_studio_lifecycle_config#tags_all TfStudioLifecycleConfig#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

