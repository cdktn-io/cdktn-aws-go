package awsamp


// Experimental.
type AwsPrometheusAnomalyDetector_RandomCutForestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#query AwsPrometheusAnomalyDetector#query}.
	// Experimental.
	Query *string `field:"required" json:"query" yaml:"query"`
	// ignore_near_expected_from_above block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#ignore_near_expected_from_above AwsPrometheusAnomalyDetector#ignore_near_expected_from_above}
	// Experimental.
	IgnoreNearExpectedFromAbove interface{} `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// ignore_near_expected_from_below block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#ignore_near_expected_from_below AwsPrometheusAnomalyDetector#ignore_near_expected_from_below}
	// Experimental.
	IgnoreNearExpectedFromBelow interface{} `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#sample_size AwsPrometheusAnomalyDetector#sample_size}.
	// Experimental.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#shingle_size AwsPrometheusAnomalyDetector#shingle_size}.
	// Experimental.
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

