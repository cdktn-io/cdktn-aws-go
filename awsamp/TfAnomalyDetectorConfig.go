package awsamp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAnomalyDetectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#alias TfAnomalyDetector#alias}.
	// Experimental.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#workspace_id TfAnomalyDetector#workspace_id}.
	// Experimental.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#configuration TfAnomalyDetector#configuration}
	// Experimental.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#evaluation_interval_in_seconds TfAnomalyDetector#evaluation_interval_in_seconds}.
	// Experimental.
	EvaluationIntervalInSeconds *float64 `field:"optional" json:"evaluationIntervalInSeconds" yaml:"evaluationIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#labels TfAnomalyDetector#labels}.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// missing_data_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#missing_data_action TfAnomalyDetector#missing_data_action}
	// Experimental.
	MissingDataAction interface{} `field:"optional" json:"missingDataAction" yaml:"missingDataAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#region TfAnomalyDetector#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#tags TfAnomalyDetector#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#timeouts TfAnomalyDetector#timeouts}
	// Experimental.
	Timeouts *TfAnomalyDetector_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

