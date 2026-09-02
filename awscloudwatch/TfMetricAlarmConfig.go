package awscloudwatch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMetricAlarmConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#alarm_name TfMetricAlarm#alarm_name}.
	// Experimental.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#actions_enabled TfMetricAlarm#actions_enabled}.
	// Experimental.
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#alarm_actions TfMetricAlarm#alarm_actions}.
	// Experimental.
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#alarm_description TfMetricAlarm#alarm_description}.
	// Experimental.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#comparison_operator TfMetricAlarm#comparison_operator}.
	// Experimental.
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#datapoints_to_alarm TfMetricAlarm#datapoints_to_alarm}.
	// Experimental.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#dimensions TfMetricAlarm#dimensions}.
	// Experimental.
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#evaluate_low_sample_count_percentiles TfMetricAlarm#evaluate_low_sample_count_percentiles}.
	// Experimental.
	EvaluateLowSampleCountPercentiles *string `field:"optional" json:"evaluateLowSampleCountPercentiles" yaml:"evaluateLowSampleCountPercentiles"`
	// evaluation_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#evaluation_criteria TfMetricAlarm#evaluation_criteria}
	// Experimental.
	EvaluationCriteria *TfMetricAlarm_EvaluationCriteriaProperty `field:"optional" json:"evaluationCriteria" yaml:"evaluationCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#evaluation_interval TfMetricAlarm#evaluation_interval}.
	// Experimental.
	EvaluationInterval *float64 `field:"optional" json:"evaluationInterval" yaml:"evaluationInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#evaluation_periods TfMetricAlarm#evaluation_periods}.
	// Experimental.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#extended_statistic TfMetricAlarm#extended_statistic}.
	// Experimental.
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#id TfMetricAlarm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#insufficient_data_actions TfMetricAlarm#insufficient_data_actions}.
	// Experimental.
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#metric_name TfMetricAlarm#metric_name}.
	// Experimental.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#metric_query TfMetricAlarm#metric_query}
	// Experimental.
	MetricQuery interface{} `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#namespace TfMetricAlarm#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#ok_actions TfMetricAlarm#ok_actions}.
	// Experimental.
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#period TfMetricAlarm#period}.
	// Experimental.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#region TfMetricAlarm#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#statistic TfMetricAlarm#statistic}.
	// Experimental.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#tags TfMetricAlarm#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#tags_all TfMetricAlarm#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#threshold TfMetricAlarm#threshold}.
	// Experimental.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#threshold_metric_id TfMetricAlarm#threshold_metric_id}.
	// Experimental.
	ThresholdMetricId *string `field:"optional" json:"thresholdMetricId" yaml:"thresholdMetricId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#treat_missing_data TfMetricAlarm#treat_missing_data}.
	// Experimental.
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm#unit TfMetricAlarm#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

