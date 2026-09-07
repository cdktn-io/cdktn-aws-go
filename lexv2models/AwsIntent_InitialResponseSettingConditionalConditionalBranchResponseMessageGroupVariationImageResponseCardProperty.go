package lexv2models


// Experimental.
type AwsIntent_InitialResponseSettingConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#title AwsIntent#title}.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// button block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#button AwsIntent#button}
	// Experimental.
	Button interface{} `field:"optional" json:"button" yaml:"button"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#image_url AwsIntent#image_url}.
	// Experimental.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#subtitle AwsIntent#subtitle}.
	// Experimental.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

