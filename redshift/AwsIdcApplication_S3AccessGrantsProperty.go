package redshift


// Experimental.
type AwsIdcApplication_S3AccessGrantsProperty struct {
	// read_write_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#read_write_access AwsIdcApplication#read_write_access}
	// Experimental.
	ReadWriteAccess interface{} `field:"optional" json:"readWriteAccess" yaml:"readWriteAccess"`
}

