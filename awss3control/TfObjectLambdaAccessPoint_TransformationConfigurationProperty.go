package awss3control


// Experimental.
type TfObjectLambdaAccessPoint_TransformationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#actions TfObjectLambdaAccessPoint#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// content_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#content_transformation TfObjectLambdaAccessPoint#content_transformation}
	// Experimental.
	ContentTransformation *TfObjectLambdaAccessPoint_ContentTransformationProperty `field:"required" json:"contentTransformation" yaml:"contentTransformation"`
}

