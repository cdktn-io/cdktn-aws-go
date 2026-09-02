package awslexv2models


// Experimental.
type TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroupVariationImageResponseCardProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#title TfSlot#title}.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// button block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#button TfSlot#button}
	// Experimental.
	Button interface{} `field:"optional" json:"button" yaml:"button"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#image_url TfSlot#image_url}.
	// Experimental.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#subtitle TfSlot#subtitle}.
	// Experimental.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

