package s3control


// Experimental.
type AwsObjectLambdaAccessPoint_AwsLambdaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#function_arn AwsObjectLambdaAccessPoint#function_arn}.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#function_payload AwsObjectLambdaAccessPoint#function_payload}.
	// Experimental.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

