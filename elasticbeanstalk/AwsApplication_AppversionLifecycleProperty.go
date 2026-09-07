package elasticbeanstalk


// Experimental.
type AwsApplication_AppversionLifecycleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_application#service_role AwsApplication#service_role}.
	// Experimental.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_application#delete_source_from_s3 AwsApplication#delete_source_from_s3}.
	// Experimental.
	DeleteSourceFromS3 interface{} `field:"optional" json:"deleteSourceFromS3" yaml:"deleteSourceFromS3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_application#max_age_in_days AwsApplication#max_age_in_days}.
	// Experimental.
	MaxAgeInDays *float64 `field:"optional" json:"maxAgeInDays" yaml:"maxAgeInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastic_beanstalk_application#max_count AwsApplication#max_count}.
	// Experimental.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
}

