package awsresiliencehubv2


// Experimental.
type TfInputSource_EksProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#cluster_arn TfInputSource#cluster_arn}.
	// Experimental.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#namespaces TfInputSource#namespaces}.
	// Experimental.
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

