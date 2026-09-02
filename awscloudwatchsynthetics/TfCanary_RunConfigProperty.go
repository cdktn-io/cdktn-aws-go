package awscloudwatchsynthetics


// Experimental.
type TfCanary_RunConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#active_tracing TfCanary#active_tracing}.
	// Experimental.
	ActiveTracing interface{} `field:"optional" json:"activeTracing" yaml:"activeTracing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#environment_variables TfCanary#environment_variables}.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#ephemeral_storage TfCanary#ephemeral_storage}.
	// Experimental.
	EphemeralStorage *float64 `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#memory_in_mb TfCanary#memory_in_mb}.
	// Experimental.
	MemoryInMb *float64 `field:"optional" json:"memoryInMb" yaml:"memoryInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#timeout_in_seconds TfCanary#timeout_in_seconds}.
	// Experimental.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

