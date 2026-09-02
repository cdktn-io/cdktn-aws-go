package awsec2


// Experimental.
type TfInstance_CpuOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#amd_sev_snp TfInstance#amd_sev_snp}.
	// Experimental.
	AmdSevSnp *string `field:"optional" json:"amdSevSnp" yaml:"amdSevSnp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#core_count TfInstance#core_count}.
	// Experimental.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#nested_virtualization TfInstance#nested_virtualization}.
	// Experimental.
	NestedVirtualization *string `field:"optional" json:"nestedVirtualization" yaml:"nestedVirtualization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#threads_per_core TfInstance#threads_per_core}.
	// Experimental.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

