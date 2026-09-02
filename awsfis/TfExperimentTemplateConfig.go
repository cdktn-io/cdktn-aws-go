package awsfis

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfExperimentTemplateConfig struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#action TfExperimentTemplate#action}
	// Experimental.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#description TfExperimentTemplate#description}.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#role_arn TfExperimentTemplate#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// stop_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#stop_condition TfExperimentTemplate#stop_condition}
	// Experimental.
	StopCondition interface{} `field:"required" json:"stopCondition" yaml:"stopCondition"`
	// experiment_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#experiment_options TfExperimentTemplate#experiment_options}
	// Experimental.
	ExperimentOptions *TfExperimentTemplate_ExperimentOptionsProperty `field:"optional" json:"experimentOptions" yaml:"experimentOptions"`
	// experiment_report_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#experiment_report_configuration TfExperimentTemplate#experiment_report_configuration}
	// Experimental.
	ExperimentReportConfiguration *TfExperimentTemplate_ExperimentReportConfigurationProperty `field:"optional" json:"experimentReportConfiguration" yaml:"experimentReportConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#id TfExperimentTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#log_configuration TfExperimentTemplate#log_configuration}
	// Experimental.
	LogConfiguration *TfExperimentTemplate_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#region TfExperimentTemplate#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#tags TfExperimentTemplate#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#tags_all TfExperimentTemplate#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#target TfExperimentTemplate#target}
	// Experimental.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#timeouts TfExperimentTemplate#timeouts}
	// Experimental.
	Timeouts *TfExperimentTemplate_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

