package awsredshift


// Experimental.
type TfIdcApplication_S3AccessGrantsProperty struct {
	// read_write_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#read_write_access TfIdcApplication#read_write_access}
	// Experimental.
	ReadWriteAccess interface{} `field:"optional" json:"readWriteAccess" yaml:"readWriteAccess"`
}

