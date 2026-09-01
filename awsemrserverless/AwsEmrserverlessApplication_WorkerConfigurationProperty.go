package awsemrserverless


// Experimental.
type AwsEmrserverlessApplication_WorkerConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#cpu AwsEmrserverlessApplication#cpu}.
	// Experimental.
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#memory AwsEmrserverlessApplication#memory}.
	// Experimental.
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#disk AwsEmrserverlessApplication#disk}.
	// Experimental.
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

