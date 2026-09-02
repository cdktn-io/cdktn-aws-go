package awsecs


// Experimental.
type TfTaskDefinition_RuntimePlatformProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#cpu_architecture TfTaskDefinition#cpu_architecture}.
	// Experimental.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#operating_system_family TfTaskDefinition#operating_system_family}.
	// Experimental.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

