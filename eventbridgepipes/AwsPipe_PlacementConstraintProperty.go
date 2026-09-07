package eventbridgepipes


// Experimental.
type AwsPipe_PlacementConstraintProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#expression AwsPipe#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#type AwsPipe#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

