package awsdatasync


// Experimental.
type AwsDatasyncLocationHdfs_NameNodeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#hostname AwsDatasyncLocationHdfs#hostname}.
	// Experimental.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#port AwsDatasyncLocationHdfs#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

