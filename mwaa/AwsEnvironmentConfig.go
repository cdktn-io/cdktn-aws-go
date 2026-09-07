package mwaa

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#dag_s3_path AwsEnvironment#dag_s3_path}.
	// Experimental.
	DagS3Path *string `field:"required" json:"dagS3Path" yaml:"dagS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#execution_role_arn AwsEnvironment#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#name AwsEnvironment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#network_configuration AwsEnvironment#network_configuration}
	// Experimental.
	NetworkConfiguration *AwsEnvironment_NetworkConfigurationProperty `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#source_bucket_arn AwsEnvironment#source_bucket_arn}.
	// Experimental.
	SourceBucketArn *string `field:"required" json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#airflow_configuration_options AwsEnvironment#airflow_configuration_options}.
	// Experimental.
	AirflowConfigurationOptions *map[string]*string `field:"optional" json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#airflow_version AwsEnvironment#airflow_version}.
	// Experimental.
	AirflowVersion *string `field:"optional" json:"airflowVersion" yaml:"airflowVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#endpoint_management AwsEnvironment#endpoint_management}.
	// Experimental.
	EndpointManagement *string `field:"optional" json:"endpointManagement" yaml:"endpointManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#environment_class AwsEnvironment#environment_class}.
	// Experimental.
	EnvironmentClass *string `field:"optional" json:"environmentClass" yaml:"environmentClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#id AwsEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#kms_key AwsEnvironment#kms_key}.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#logging_configuration AwsEnvironment#logging_configuration}
	// Experimental.
	LoggingConfiguration *AwsEnvironment_LoggingConfigurationProperty `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#max_webservers AwsEnvironment#max_webservers}.
	// Experimental.
	MaxWebservers *float64 `field:"optional" json:"maxWebservers" yaml:"maxWebservers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#max_workers AwsEnvironment#max_workers}.
	// Experimental.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#min_webservers AwsEnvironment#min_webservers}.
	// Experimental.
	MinWebservers *float64 `field:"optional" json:"minWebservers" yaml:"minWebservers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#min_workers AwsEnvironment#min_workers}.
	// Experimental.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#plugins_s3_object_version AwsEnvironment#plugins_s3_object_version}.
	// Experimental.
	PluginsS3ObjectVersion *string `field:"optional" json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#plugins_s3_path AwsEnvironment#plugins_s3_path}.
	// Experimental.
	PluginsS3Path *string `field:"optional" json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#region AwsEnvironment#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#requirements_s3_object_version AwsEnvironment#requirements_s3_object_version}.
	// Experimental.
	RequirementsS3ObjectVersion *string `field:"optional" json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#requirements_s3_path AwsEnvironment#requirements_s3_path}.
	// Experimental.
	RequirementsS3Path *string `field:"optional" json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#schedulers AwsEnvironment#schedulers}.
	// Experimental.
	Schedulers *float64 `field:"optional" json:"schedulers" yaml:"schedulers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#startup_script_s3_object_version AwsEnvironment#startup_script_s3_object_version}.
	// Experimental.
	StartupScriptS3ObjectVersion *string `field:"optional" json:"startupScriptS3ObjectVersion" yaml:"startupScriptS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#startup_script_s3_path AwsEnvironment#startup_script_s3_path}.
	// Experimental.
	StartupScriptS3Path *string `field:"optional" json:"startupScriptS3Path" yaml:"startupScriptS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#tags AwsEnvironment#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#tags_all AwsEnvironment#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#timeouts AwsEnvironment#timeouts}
	// Experimental.
	Timeouts *AwsEnvironment_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#webserver_access_mode AwsEnvironment#webserver_access_mode}.
	// Experimental.
	WebserverAccessMode *string `field:"optional" json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#weekly_maintenance_window_start AwsEnvironment#weekly_maintenance_window_start}.
	// Experimental.
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#worker_replacement_strategy AwsEnvironment#worker_replacement_strategy}.
	// Experimental.
	WorkerReplacementStrategy *string `field:"optional" json:"workerReplacementStrategy" yaml:"workerReplacementStrategy"`
}

