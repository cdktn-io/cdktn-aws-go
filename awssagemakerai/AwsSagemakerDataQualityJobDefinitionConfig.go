package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerDataQualityJobDefinitionConfig struct {
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
	// data_quality_app_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_app_specification AwsSagemakerDataQualityJobDefinition#data_quality_app_specification}
	// Experimental.
	DataQualityAppSpecification *AwsSagemakerDataQualityJobDefinition_DataQualityAppSpecificationProperty `field:"required" json:"dataQualityAppSpecification" yaml:"dataQualityAppSpecification"`
	// data_quality_job_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_job_input AwsSagemakerDataQualityJobDefinition#data_quality_job_input}
	// Experimental.
	DataQualityJobInput *AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty `field:"required" json:"dataQualityJobInput" yaml:"dataQualityJobInput"`
	// data_quality_job_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_job_output_config AwsSagemakerDataQualityJobDefinition#data_quality_job_output_config}
	// Experimental.
	DataQualityJobOutputConfig *AwsSagemakerDataQualityJobDefinition_DataQualityJobOutputConfigProperty `field:"required" json:"dataQualityJobOutputConfig" yaml:"dataQualityJobOutputConfig"`
	// job_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#job_resources AwsSagemakerDataQualityJobDefinition#job_resources}
	// Experimental.
	JobResources *AwsSagemakerDataQualityJobDefinition_JobResourcesProperty `field:"required" json:"jobResources" yaml:"jobResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#role_arn AwsSagemakerDataQualityJobDefinition#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// data_quality_baseline_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_baseline_config AwsSagemakerDataQualityJobDefinition#data_quality_baseline_config}
	// Experimental.
	DataQualityBaselineConfig *AwsSagemakerDataQualityJobDefinition_DataQualityBaselineConfigProperty `field:"optional" json:"dataQualityBaselineConfig" yaml:"dataQualityBaselineConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#id AwsSagemakerDataQualityJobDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#name AwsSagemakerDataQualityJobDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#network_config AwsSagemakerDataQualityJobDefinition#network_config}
	// Experimental.
	NetworkConfig *AwsSagemakerDataQualityJobDefinition_NetworkConfigProperty `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#region AwsSagemakerDataQualityJobDefinition#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#stopping_condition AwsSagemakerDataQualityJobDefinition#stopping_condition}
	// Experimental.
	StoppingCondition *AwsSagemakerDataQualityJobDefinition_StoppingConditionProperty `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#tags AwsSagemakerDataQualityJobDefinition#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#tags_all AwsSagemakerDataQualityJobDefinition#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

