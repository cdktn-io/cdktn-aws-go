package lexv2models


// Experimental.
type AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageImageResponseCardProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#title AwsSlot#title}.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// button block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#button AwsSlot#button}
	// Experimental.
	Button interface{} `field:"optional" json:"button" yaml:"button"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#image_url AwsSlot#image_url}.
	// Experimental.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#subtitle AwsSlot#subtitle}.
	// Experimental.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

