package awsemrcontainers


// Experimental.
type TfJobTemplate_JobTemplateDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#execution_role_arn TfJobTemplate#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#job_driver TfJobTemplate#job_driver}
	// Experimental.
	JobDriver *TfJobTemplate_JobDriverProperty `field:"required" json:"jobDriver" yaml:"jobDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#release_label TfJobTemplate#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// configuration_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#configuration_overrides TfJobTemplate#configuration_overrides}
	// Experimental.
	ConfigurationOverrides *TfJobTemplate_ConfigurationOverridesProperty `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#job_tags TfJobTemplate#job_tags}.
	// Experimental.
	JobTags *map[string]*string `field:"optional" json:"jobTags" yaml:"jobTags"`
}

