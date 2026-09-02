package awsemrserverless


// Experimental.
type TfApplication_MaximumCapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#cpu TfApplication#cpu}.
	// Experimental.
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#memory TfApplication#memory}.
	// Experimental.
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#disk TfApplication#disk}.
	// Experimental.
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

