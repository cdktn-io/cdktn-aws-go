package eventbridgepipes


// Experimental.
type AwsPipe_EphemeralStorageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#size_in_gib AwsPipe#size_in_gib}.
	// Experimental.
	SizeInGib *float64 `field:"required" json:"sizeInGib" yaml:"sizeInGib"`
}

