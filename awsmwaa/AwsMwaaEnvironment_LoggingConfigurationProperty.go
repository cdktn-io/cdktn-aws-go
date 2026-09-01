package awsmwaa


// Experimental.
type AwsMwaaEnvironment_LoggingConfigurationProperty struct {
	// dag_processing_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#dag_processing_logs AwsMwaaEnvironment#dag_processing_logs}
	// Experimental.
	DagProcessingLogs *AwsMwaaEnvironment_DagProcessingLogsProperty `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// scheduler_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#scheduler_logs AwsMwaaEnvironment#scheduler_logs}
	// Experimental.
	SchedulerLogs *AwsMwaaEnvironment_SchedulerLogsProperty `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// task_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#task_logs AwsMwaaEnvironment#task_logs}
	// Experimental.
	TaskLogs *AwsMwaaEnvironment_TaskLogsProperty `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// webserver_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#webserver_logs AwsMwaaEnvironment#webserver_logs}
	// Experimental.
	WebserverLogs *AwsMwaaEnvironment_WebserverLogsProperty `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// worker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#worker_logs AwsMwaaEnvironment#worker_logs}
	// Experimental.
	WorkerLogs *AwsMwaaEnvironment_WorkerLogsProperty `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

