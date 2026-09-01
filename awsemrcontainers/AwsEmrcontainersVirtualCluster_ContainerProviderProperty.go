package awsemrcontainers


// Experimental.
type AwsEmrcontainersVirtualCluster_ContainerProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_virtual_cluster#id AwsEmrcontainersVirtualCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_virtual_cluster#info AwsEmrcontainersVirtualCluster#info}
	// Experimental.
	Info *AwsEmrcontainersVirtualCluster_InfoProperty `field:"required" json:"info" yaml:"info"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_virtual_cluster#type AwsEmrcontainersVirtualCluster#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

