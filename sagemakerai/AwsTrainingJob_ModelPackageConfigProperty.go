package sagemakerai


// Experimental.
type AwsTrainingJob_ModelPackageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#model_package_group_arn AwsTrainingJob#model_package_group_arn}.
	// Experimental.
	ModelPackageGroupArn *string `field:"required" json:"modelPackageGroupArn" yaml:"modelPackageGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#source_model_package_arn AwsTrainingJob#source_model_package_arn}.
	// Experimental.
	SourceModelPackageArn *string `field:"optional" json:"sourceModelPackageArn" yaml:"sourceModelPackageArn"`
}

