package awsmwaa


// Experimental.
type TfEnvironment_LoggingConfigurationProperty struct {
	// dag_processing_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#dag_processing_logs TfEnvironment#dag_processing_logs}
	// Experimental.
	DagProcessingLogs *TfEnvironment_DagProcessingLogsProperty `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// scheduler_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#scheduler_logs TfEnvironment#scheduler_logs}
	// Experimental.
	SchedulerLogs *TfEnvironment_SchedulerLogsProperty `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// task_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#task_logs TfEnvironment#task_logs}
	// Experimental.
	TaskLogs *TfEnvironment_TaskLogsProperty `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// webserver_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#webserver_logs TfEnvironment#webserver_logs}
	// Experimental.
	WebserverLogs *TfEnvironment_WebserverLogsProperty `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// worker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#worker_logs TfEnvironment#worker_logs}
	// Experimental.
	WorkerLogs *TfEnvironment_WorkerLogsProperty `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

