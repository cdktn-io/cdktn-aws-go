package awsecs


// Experimental.
type AwsEcsTaskDefinition_RuntimePlatformProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#cpu_architecture AwsEcsTaskDefinition#cpu_architecture}.
	// Experimental.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#operating_system_family AwsEcsTaskDefinition#operating_system_family}.
	// Experimental.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

