package mwaa


// Experimental.
type AwsEnvironment_LoggingConfigurationProperty struct {
	// dag_processing_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#dag_processing_logs AwsEnvironment#dag_processing_logs}
	// Experimental.
	DagProcessingLogs *AwsEnvironment_DagProcessingLogsProperty `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// scheduler_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#scheduler_logs AwsEnvironment#scheduler_logs}
	// Experimental.
	SchedulerLogs *AwsEnvironment_SchedulerLogsProperty `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// task_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#task_logs AwsEnvironment#task_logs}
	// Experimental.
	TaskLogs *AwsEnvironment_TaskLogsProperty `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// webserver_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#webserver_logs AwsEnvironment#webserver_logs}
	// Experimental.
	WebserverLogs *AwsEnvironment_WebserverLogsProperty `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// worker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#worker_logs AwsEnvironment#worker_logs}
	// Experimental.
	WorkerLogs *AwsEnvironment_WorkerLogsProperty `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

