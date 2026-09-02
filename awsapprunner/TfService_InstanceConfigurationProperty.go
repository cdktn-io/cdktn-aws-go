package awsapprunner


// Experimental.
type TfService_InstanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#cpu TfService#cpu}.
	// Experimental.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#instance_role_arn TfService#instance_role_arn}.
	// Experimental.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#memory TfService#memory}.
	// Experimental.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

