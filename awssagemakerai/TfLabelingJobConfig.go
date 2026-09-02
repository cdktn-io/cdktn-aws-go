package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLabelingJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#label_attribute_name TfLabelingJob#label_attribute_name}.
	// Experimental.
	LabelAttributeName *string `field:"required" json:"labelAttributeName" yaml:"labelAttributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#labeling_job_name TfLabelingJob#labeling_job_name}.
	// Experimental.
	LabelingJobName *string `field:"required" json:"labelingJobName" yaml:"labelingJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#role_arn TfLabelingJob#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// human_task_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#human_task_config TfLabelingJob#human_task_config}
	// Experimental.
	HumanTaskConfig interface{} `field:"optional" json:"humanTaskConfig" yaml:"humanTaskConfig"`
	// input_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#input_config TfLabelingJob#input_config}
	// Experimental.
	InputConfig interface{} `field:"optional" json:"inputConfig" yaml:"inputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#label_category_config_s3_uri TfLabelingJob#label_category_config_s3_uri}.
	// Experimental.
	LabelCategoryConfigS3Uri *string `field:"optional" json:"labelCategoryConfigS3Uri" yaml:"labelCategoryConfigS3Uri"`
	// labeling_job_algorithms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#labeling_job_algorithms_config TfLabelingJob#labeling_job_algorithms_config}
	// Experimental.
	LabelingJobAlgorithmsConfig interface{} `field:"optional" json:"labelingJobAlgorithmsConfig" yaml:"labelingJobAlgorithmsConfig"`
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#output_config TfLabelingJob#output_config}
	// Experimental.
	OutputConfig interface{} `field:"optional" json:"outputConfig" yaml:"outputConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#region TfLabelingJob#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#stopping_conditions TfLabelingJob#stopping_conditions}.
	// Experimental.
	StoppingConditions interface{} `field:"optional" json:"stoppingConditions" yaml:"stoppingConditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#tags TfLabelingJob#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

