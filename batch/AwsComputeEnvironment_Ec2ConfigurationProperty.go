package batch


// Experimental.
type AwsComputeEnvironment_Ec2ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#image_id_override AwsComputeEnvironment#image_id_override}.
	// Experimental.
	ImageIdOverride *string `field:"optional" json:"imageIdOverride" yaml:"imageIdOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#image_kubernetes_version AwsComputeEnvironment#image_kubernetes_version}.
	// Experimental.
	ImageKubernetesVersion *string `field:"optional" json:"imageKubernetesVersion" yaml:"imageKubernetesVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#image_type AwsComputeEnvironment#image_type}.
	// Experimental.
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
}

