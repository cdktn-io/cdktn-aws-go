package awstimestreamforinfluxdb


// Experimental.
type TfDbCluster_LogDeliveryConfigurationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_cluster#s3_configuration TfDbCluster#s3_configuration}
	// Experimental.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

