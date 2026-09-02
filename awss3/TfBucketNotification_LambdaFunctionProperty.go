package awss3


// Experimental.
type TfBucketNotification_LambdaFunctionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#events TfBucketNotification#events}.
	// Experimental.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#filter_prefix TfBucketNotification#filter_prefix}.
	// Experimental.
	FilterPrefix *string `field:"optional" json:"filterPrefix" yaml:"filterPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#filter_suffix TfBucketNotification#filter_suffix}.
	// Experimental.
	FilterSuffix *string `field:"optional" json:"filterSuffix" yaml:"filterSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#id TfBucketNotification#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#lambda_function_arn TfBucketNotification#lambda_function_arn}.
	// Experimental.
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

