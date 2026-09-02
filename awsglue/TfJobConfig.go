package awsglue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobConfig struct {
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
	// command block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#command TfJob#command}
	// Experimental.
	Command *TfJob_CommandProperty `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#name TfJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#role_arn TfJob#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#connections TfJob#connections}.
	// Experimental.
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#default_arguments TfJob#default_arguments}.
	// Experimental.
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#description TfJob#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#execution_class TfJob#execution_class}.
	// Experimental.
	ExecutionClass *string `field:"optional" json:"executionClass" yaml:"executionClass"`
	// execution_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#execution_property TfJob#execution_property}
	// Experimental.
	ExecutionProperty *TfJob_ExecutionPropertyProperty `field:"optional" json:"executionProperty" yaml:"executionProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#glue_version TfJob#glue_version}.
	// Experimental.
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#id TfJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#job_mode TfJob#job_mode}.
	// Experimental.
	JobMode *string `field:"optional" json:"jobMode" yaml:"jobMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#job_run_queuing_enabled TfJob#job_run_queuing_enabled}.
	// Experimental.
	JobRunQueuingEnabled interface{} `field:"optional" json:"jobRunQueuingEnabled" yaml:"jobRunQueuingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#maintenance_window TfJob#maintenance_window}.
	// Experimental.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#max_capacity TfJob#max_capacity}.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#max_retries TfJob#max_retries}.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#non_overridable_arguments TfJob#non_overridable_arguments}.
	// Experimental.
	NonOverridableArguments *map[string]*string `field:"optional" json:"nonOverridableArguments" yaml:"nonOverridableArguments"`
	// notification_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#notification_property TfJob#notification_property}
	// Experimental.
	NotificationProperty *TfJob_NotificationPropertyProperty `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#number_of_workers TfJob#number_of_workers}.
	// Experimental.
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#region TfJob#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#security_configuration TfJob#security_configuration}.
	// Experimental.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// source_control_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#source_control_details TfJob#source_control_details}
	// Experimental.
	SourceControlDetails *TfJob_SourceControlDetailsProperty `field:"optional" json:"sourceControlDetails" yaml:"sourceControlDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#tags TfJob#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#tags_all TfJob#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#timeout TfJob#timeout}.
	// Experimental.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#worker_type TfJob#worker_type}.
	// Experimental.
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

