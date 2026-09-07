package sagemakerai


// Experimental.
type AwsAppImageConfig_KernelGatewayImageConfigFileSystemConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#default_gid AwsAppImageConfig#default_gid}.
	// Experimental.
	DefaultGid *float64 `field:"optional" json:"defaultGid" yaml:"defaultGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#default_uid AwsAppImageConfig#default_uid}.
	// Experimental.
	DefaultUid *float64 `field:"optional" json:"defaultUid" yaml:"defaultUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#mount_path AwsAppImageConfig#mount_path}.
	// Experimental.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

