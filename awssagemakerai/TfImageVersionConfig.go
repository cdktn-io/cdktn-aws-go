package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfImageVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#base_image TfImageVersion#base_image}.
	// Experimental.
	BaseImage *string `field:"required" json:"baseImage" yaml:"baseImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#image_name TfImageVersion#image_name}.
	// Experimental.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#aliases TfImageVersion#aliases}.
	// Experimental.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#horovod TfImageVersion#horovod}.
	// Experimental.
	Horovod interface{} `field:"optional" json:"horovod" yaml:"horovod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#id TfImageVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#job_type TfImageVersion#job_type}.
	// Experimental.
	JobType *string `field:"optional" json:"jobType" yaml:"jobType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#ml_framework TfImageVersion#ml_framework}.
	// Experimental.
	MlFramework *string `field:"optional" json:"mlFramework" yaml:"mlFramework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#processor TfImageVersion#processor}.
	// Experimental.
	Processor *string `field:"optional" json:"processor" yaml:"processor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#programming_lang TfImageVersion#programming_lang}.
	// Experimental.
	ProgrammingLang *string `field:"optional" json:"programmingLang" yaml:"programmingLang"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#region TfImageVersion#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#release_notes TfImageVersion#release_notes}.
	// Experimental.
	ReleaseNotes *string `field:"optional" json:"releaseNotes" yaml:"releaseNotes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_image_version#vendor_guidance TfImageVersion#vendor_guidance}.
	// Experimental.
	VendorGuidance *string `field:"optional" json:"vendorGuidance" yaml:"vendorGuidance"`
}

