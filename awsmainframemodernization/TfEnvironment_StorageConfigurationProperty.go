package awsmainframemodernization


// Experimental.
type TfEnvironment_StorageConfigurationProperty struct {
	// efs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/m2_environment#efs TfEnvironment#efs}
	// Experimental.
	Efs interface{} `field:"optional" json:"efs" yaml:"efs"`
	// fsx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/m2_environment#fsx TfEnvironment#fsx}
	// Experimental.
	Fsx interface{} `field:"optional" json:"fsx" yaml:"fsx"`
}

