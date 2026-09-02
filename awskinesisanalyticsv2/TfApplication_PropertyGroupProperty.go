package awskinesisanalyticsv2


// Experimental.
type TfApplication_PropertyGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#property_group_id TfApplication#property_group_id}.
	// Experimental.
	PropertyGroupId *string `field:"required" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#property_map TfApplication#property_map}.
	// Experimental.
	PropertyMap *map[string]*string `field:"required" json:"propertyMap" yaml:"propertyMap"`
}

