package awsemrserverless


// Experimental.
type AwsEmrserverlessApplication_InitialCapacityConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#worker_count AwsEmrserverlessApplication#worker_count}.
	// Experimental.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// worker_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#worker_configuration AwsEmrserverlessApplication#worker_configuration}
	// Experimental.
	WorkerConfiguration *AwsEmrserverlessApplication_WorkerConfigurationProperty `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

