package awsvpclattice


// Experimental.
type TfTargetGroupAttachment_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group_attachment#id TfTargetGroupAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group_attachment#port TfTargetGroupAttachment#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

