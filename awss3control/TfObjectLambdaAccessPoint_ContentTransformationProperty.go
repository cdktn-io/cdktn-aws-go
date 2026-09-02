package awss3control


// Experimental.
type TfObjectLambdaAccessPoint_ContentTransformationProperty struct {
	// aws_lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#aws_lambda TfObjectLambdaAccessPoint#aws_lambda}
	// Experimental.
	AwsLambda *TfObjectLambdaAccessPoint_AwsLambdaProperty `field:"required" json:"awsLambda" yaml:"awsLambda"`
}

