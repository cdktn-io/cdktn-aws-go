package awss3control


// Experimental.
type AwsS3ControlObjectLambdaAccessPoint_AwsLambdaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#function_arn AwsS3ControlObjectLambdaAccessPoint#function_arn}.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#function_payload AwsS3ControlObjectLambdaAccessPoint#function_payload}.
	// Experimental.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

