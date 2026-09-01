package awsemrcontainers


// Experimental.
type AwsEmrcontainersJobTemplate_JobTemplateDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#execution_role_arn AwsEmrcontainersJobTemplate#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#job_driver AwsEmrcontainersJobTemplate#job_driver}
	// Experimental.
	JobDriver *AwsEmrcontainersJobTemplate_JobDriverProperty `field:"required" json:"jobDriver" yaml:"jobDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#release_label AwsEmrcontainersJobTemplate#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// configuration_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#configuration_overrides AwsEmrcontainersJobTemplate#configuration_overrides}
	// Experimental.
	ConfigurationOverrides *AwsEmrcontainersJobTemplate_ConfigurationOverridesProperty `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#job_tags AwsEmrcontainersJobTemplate#job_tags}.
	// Experimental.
	JobTags *map[string]*string `field:"optional" json:"jobTags" yaml:"jobTags"`
}

