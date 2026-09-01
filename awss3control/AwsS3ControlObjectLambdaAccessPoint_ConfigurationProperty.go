package awss3control


// Experimental.
type AwsS3ControlObjectLambdaAccessPoint_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#supporting_access_point AwsS3ControlObjectLambdaAccessPoint#supporting_access_point}.
	// Experimental.
	SupportingAccessPoint *string `field:"required" json:"supportingAccessPoint" yaml:"supportingAccessPoint"`
	// transformation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#transformation_configuration AwsS3ControlObjectLambdaAccessPoint#transformation_configuration}
	// Experimental.
	TransformationConfiguration interface{} `field:"required" json:"transformationConfiguration" yaml:"transformationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#allowed_features AwsS3ControlObjectLambdaAccessPoint#allowed_features}.
	// Experimental.
	AllowedFeatures *[]*string `field:"optional" json:"allowedFeatures" yaml:"allowedFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#cloud_watch_metrics_enabled AwsS3ControlObjectLambdaAccessPoint#cloud_watch_metrics_enabled}.
	// Experimental.
	CloudWatchMetricsEnabled interface{} `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
}

