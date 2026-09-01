package awsec2


// Experimental.
type AwsEc2AllowedImagesSettings_ImageCriterionProperty struct {
	// creation_date_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#creation_date_condition AwsEc2AllowedImagesSettings#creation_date_condition}
	// Experimental.
	CreationDateCondition interface{} `field:"optional" json:"creationDateCondition" yaml:"creationDateCondition"`
	// deprecation_time_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#deprecation_time_condition AwsEc2AllowedImagesSettings#deprecation_time_condition}
	// Experimental.
	DeprecationTimeCondition interface{} `field:"optional" json:"deprecationTimeCondition" yaml:"deprecationTimeCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#image_names AwsEc2AllowedImagesSettings#image_names}.
	// Experimental.
	ImageNames *[]*string `field:"optional" json:"imageNames" yaml:"imageNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#image_providers AwsEc2AllowedImagesSettings#image_providers}.
	// Experimental.
	ImageProviders *[]*string `field:"optional" json:"imageProviders" yaml:"imageProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#marketplace_product_codes AwsEc2AllowedImagesSettings#marketplace_product_codes}.
	// Experimental.
	MarketplaceProductCodes *[]*string `field:"optional" json:"marketplaceProductCodes" yaml:"marketplaceProductCodes"`
}

