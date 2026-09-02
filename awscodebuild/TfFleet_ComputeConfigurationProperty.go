package awscodebuild


// Experimental.
type TfFleet_ComputeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#disk TfFleet#disk}.
	// Experimental.
	Disk *float64 `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#instance_type TfFleet#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#machine_type TfFleet#machine_type}.
	// Experimental.
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#memory TfFleet#memory}.
	// Experimental.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#vcpu TfFleet#vcpu}.
	// Experimental.
	Vcpu *float64 `field:"optional" json:"vcpu" yaml:"vcpu"`
}

