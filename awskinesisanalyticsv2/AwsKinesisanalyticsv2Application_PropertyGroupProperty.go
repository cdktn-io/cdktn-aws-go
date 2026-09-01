package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_PropertyGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#property_group_id AwsKinesisanalyticsv2Application#property_group_id}.
	// Experimental.
	PropertyGroupId *string `field:"required" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#property_map AwsKinesisanalyticsv2Application#property_map}.
	// Experimental.
	PropertyMap *map[string]*string `field:"required" json:"propertyMap" yaml:"propertyMap"`
}

