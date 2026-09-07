package ecs


// Experimental.
type AwsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#access_point_id AwsTaskDefinition#access_point_id}.
	// Experimental.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#iam AwsTaskDefinition#iam}.
	// Experimental.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

