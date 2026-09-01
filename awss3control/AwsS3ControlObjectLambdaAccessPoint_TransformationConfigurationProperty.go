package awss3control


// Experimental.
type AwsS3ControlObjectLambdaAccessPoint_TransformationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#actions AwsS3ControlObjectLambdaAccessPoint#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// content_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#content_transformation AwsS3ControlObjectLambdaAccessPoint#content_transformation}
	// Experimental.
	ContentTransformation *AwsS3ControlObjectLambdaAccessPoint_ContentTransformationProperty `field:"required" json:"contentTransformation" yaml:"contentTransformation"`
}

