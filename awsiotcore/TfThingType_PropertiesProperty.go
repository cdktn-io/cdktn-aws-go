package awsiotcore


// Experimental.
type TfThingType_PropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_type#description TfThingType#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_type#searchable_attributes TfThingType#searchable_attributes}.
	// Experimental.
	SearchableAttributes *[]*string `field:"optional" json:"searchableAttributes" yaml:"searchableAttributes"`
}

