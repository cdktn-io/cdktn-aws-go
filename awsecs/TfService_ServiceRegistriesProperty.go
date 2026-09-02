package awsecs


// Experimental.
type TfService_ServiceRegistriesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#registry_arn TfService#registry_arn}.
	// Experimental.
	RegistryArn *string `field:"required" json:"registryArn" yaml:"registryArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#container_name TfService#container_name}.
	// Experimental.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#container_port TfService#container_port}.
	// Experimental.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#port TfService#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

