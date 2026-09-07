package emrserverless


// Experimental.
type AwsApplication_MaximumCapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#cpu AwsApplication#cpu}.
	// Experimental.
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#memory AwsApplication#memory}.
	// Experimental.
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#disk AwsApplication#disk}.
	// Experimental.
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

