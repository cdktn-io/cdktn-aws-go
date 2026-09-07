package s3control


// Experimental.
type AwsObjectLambdaAccessPoint_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#supporting_access_point AwsObjectLambdaAccessPoint#supporting_access_point}.
	// Experimental.
	SupportingAccessPoint *string `field:"required" json:"supportingAccessPoint" yaml:"supportingAccessPoint"`
	// transformation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#transformation_configuration AwsObjectLambdaAccessPoint#transformation_configuration}
	// Experimental.
	TransformationConfiguration interface{} `field:"required" json:"transformationConfiguration" yaml:"transformationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#allowed_features AwsObjectLambdaAccessPoint#allowed_features}.
	// Experimental.
	AllowedFeatures *[]*string `field:"optional" json:"allowedFeatures" yaml:"allowedFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_object_lambda_access_point#cloud_watch_metrics_enabled AwsObjectLambdaAccessPoint#cloud_watch_metrics_enabled}.
	// Experimental.
	CloudWatchMetricsEnabled interface{} `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
}

