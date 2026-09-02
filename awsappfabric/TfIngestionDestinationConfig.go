package awsappfabric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIngestionDestinationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#app_bundle_arn TfIngestionDestination#app_bundle_arn}.
	// Experimental.
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#ingestion_arn TfIngestionDestination#ingestion_arn}.
	// Experimental.
	IngestionArn *string `field:"required" json:"ingestionArn" yaml:"ingestionArn"`
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#destination_configuration TfIngestionDestination#destination_configuration}
	// Experimental.
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#processing_configuration TfIngestionDestination#processing_configuration}
	// Experimental.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#region TfIngestionDestination#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#tags TfIngestionDestination#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#timeouts TfIngestionDestination#timeouts}
	// Experimental.
	Timeouts *TfIngestionDestination_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

