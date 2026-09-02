package awsglue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClassifierConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#name TfClassifier#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// csv_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#csv_classifier TfClassifier#csv_classifier}
	// Experimental.
	CsvClassifier *TfClassifier_CsvClassifierProperty `field:"optional" json:"csvClassifier" yaml:"csvClassifier"`
	// grok_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#grok_classifier TfClassifier#grok_classifier}
	// Experimental.
	GrokClassifier *TfClassifier_GrokClassifierProperty `field:"optional" json:"grokClassifier" yaml:"grokClassifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#id TfClassifier#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// json_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#json_classifier TfClassifier#json_classifier}
	// Experimental.
	JsonClassifier *TfClassifier_JsonClassifierProperty `field:"optional" json:"jsonClassifier" yaml:"jsonClassifier"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#region TfClassifier#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// xml_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#xml_classifier TfClassifier#xml_classifier}
	// Experimental.
	XmlClassifier *TfClassifier_XmlClassifierProperty `field:"optional" json:"xmlClassifier" yaml:"xmlClassifier"`
}

