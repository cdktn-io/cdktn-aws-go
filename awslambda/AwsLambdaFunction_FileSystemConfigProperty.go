package awslambda


// Experimental.
type AwsLambdaFunction_FileSystemConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#arn AwsLambdaFunction#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#local_mount_path AwsLambdaFunction#local_mount_path}.
	// Experimental.
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
}

