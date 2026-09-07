package sagemakerai


// Experimental.
type AwsAlgorithm_TransformResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_count AwsAlgorithm#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_type AwsAlgorithm#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#transform_ami_version AwsAlgorithm#transform_ami_version}.
	// Experimental.
	TransformAmiVersion *string `field:"optional" json:"transformAmiVersion" yaml:"transformAmiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#volume_kms_key_id AwsAlgorithm#volume_kms_key_id}.
	// Experimental.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

