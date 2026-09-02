package awsemr


// Experimental.
type TfCluster_HadoopJarStepProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#args TfCluster#args}.
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#jar TfCluster#jar}.
	// Experimental.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#main_class TfCluster#main_class}.
	// Experimental.
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#properties TfCluster#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

