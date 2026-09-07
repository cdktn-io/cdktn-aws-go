package lightsail


// Experimental.
type AwsContainerService_PrivateRegistryAccessProperty struct {
	// ecr_image_puller_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service#ecr_image_puller_role AwsContainerService#ecr_image_puller_role}
	// Experimental.
	EcrImagePullerRole *AwsContainerService_EcrImagePullerRoleProperty `field:"optional" json:"ecrImagePullerRole" yaml:"ecrImagePullerRole"`
}

