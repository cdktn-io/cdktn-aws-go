package emrserverless


// Experimental.
type AwsApplication_InitialCapacityConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#worker_count AwsApplication#worker_count}.
	// Experimental.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// worker_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#worker_configuration AwsApplication#worker_configuration}
	// Experimental.
	WorkerConfiguration *AwsApplication_WorkerConfigurationProperty `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

