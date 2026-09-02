package awss3control


// Experimental.
type TfObjectLambdaAccessPoint_AwsLambdaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#function_arn TfObjectLambdaAccessPoint#function_arn}.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#function_payload TfObjectLambdaAccessPoint#function_payload}.
	// Experimental.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

