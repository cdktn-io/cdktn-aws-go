package ec2


// Experimental.
type AwsSpotInstanceRequest_CpuOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#amd_sev_snp AwsSpotInstanceRequest#amd_sev_snp}.
	// Experimental.
	AmdSevSnp *string `field:"optional" json:"amdSevSnp" yaml:"amdSevSnp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#core_count AwsSpotInstanceRequest#core_count}.
	// Experimental.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#nested_virtualization AwsSpotInstanceRequest#nested_virtualization}.
	// Experimental.
	NestedVirtualization *string `field:"optional" json:"nestedVirtualization" yaml:"nestedVirtualization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#threads_per_core AwsSpotInstanceRequest#threads_per_core}.
	// Experimental.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

