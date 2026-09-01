package awsdatasync


// Experimental.
type AwsDatasyncLocationHdfs_QopConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#data_transfer_protection AwsDatasyncLocationHdfs#data_transfer_protection}.
	// Experimental.
	DataTransferProtection *string `field:"optional" json:"dataTransferProtection" yaml:"dataTransferProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs#rpc_protection AwsDatasyncLocationHdfs#rpc_protection}.
	// Experimental.
	RpcProtection *string `field:"optional" json:"rpcProtection" yaml:"rpcProtection"`
}

