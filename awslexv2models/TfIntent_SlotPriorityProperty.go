package awslexv2models


// Experimental.
type TfIntent_SlotPriorityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#priority TfIntent#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#slot_id TfIntent#slot_id}.
	// Experimental.
	SlotId *string `field:"required" json:"slotId" yaml:"slotId"`
}

