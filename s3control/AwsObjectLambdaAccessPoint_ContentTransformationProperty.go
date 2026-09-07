package s3control


// Experimental.
type AwsObjectLambdaAccessPoint_ContentTransformationProperty struct {
	// aws_lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#aws_lambda AwsObjectLambdaAccessPoint#aws_lambda}
	// Experimental.
	AwsLambda *AwsObjectLambdaAccessPoint_AwsLambdaProperty `field:"required" json:"awsLambda" yaml:"awsLambda"`
}

