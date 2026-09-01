package awsdevopsguru

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDevopsguruServiceIntegrationConfig struct {
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
	// kms_server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#kms_server_side_encryption AwsDevopsguruServiceIntegration#kms_server_side_encryption}
	// Experimental.
	KmsServerSideEncryption interface{} `field:"optional" json:"kmsServerSideEncryption" yaml:"kmsServerSideEncryption"`
	// logs_anomaly_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#logs_anomaly_detection AwsDevopsguruServiceIntegration#logs_anomaly_detection}
	// Experimental.
	LogsAnomalyDetection interface{} `field:"optional" json:"logsAnomalyDetection" yaml:"logsAnomalyDetection"`
	// ops_center block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#ops_center AwsDevopsguruServiceIntegration#ops_center}
	// Experimental.
	OpsCenter interface{} `field:"optional" json:"opsCenter" yaml:"opsCenter"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#region AwsDevopsguruServiceIntegration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

