package awssagemakerai


// Experimental.
type TfApp_ResourceSpecProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app#instance_type TfApp#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app#lifecycle_config_arn TfApp#lifecycle_config_arn}.
	// Experimental.
	LifecycleConfigArn *string `field:"optional" json:"lifecycleConfigArn" yaml:"lifecycleConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app#sagemaker_image_arn TfApp#sagemaker_image_arn}.
	// Experimental.
	SagemakerImageArn *string `field:"optional" json:"sagemakerImageArn" yaml:"sagemakerImageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app#sagemaker_image_version_alias TfApp#sagemaker_image_version_alias}.
	// Experimental.
	SagemakerImageVersionAlias *string `field:"optional" json:"sagemakerImageVersionAlias" yaml:"sagemakerImageVersionAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app#sagemaker_image_version_arn TfApp#sagemaker_image_version_arn}.
	// Experimental.
	SagemakerImageVersionArn *string `field:"optional" json:"sagemakerImageVersionArn" yaml:"sagemakerImageVersionArn"`
}

