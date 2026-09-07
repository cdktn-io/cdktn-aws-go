package ec2


// Experimental.
type AwsLaunchTemplate_CpuOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#amd_sev_snp AwsLaunchTemplate#amd_sev_snp}.
	// Experimental.
	AmdSevSnp *string `field:"optional" json:"amdSevSnp" yaml:"amdSevSnp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#core_count AwsLaunchTemplate#core_count}.
	// Experimental.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#nested_virtualization AwsLaunchTemplate#nested_virtualization}.
	// Experimental.
	NestedVirtualization *string `field:"optional" json:"nestedVirtualization" yaml:"nestedVirtualization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#threads_per_core AwsLaunchTemplate#threads_per_core}.
	// Experimental.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

