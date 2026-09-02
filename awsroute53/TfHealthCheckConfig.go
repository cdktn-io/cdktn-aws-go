package awsroute53

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHealthCheckConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#type TfHealthCheck#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#child_healthchecks TfHealthCheck#child_healthchecks}.
	// Experimental.
	ChildHealthchecks *[]*string `field:"optional" json:"childHealthchecks" yaml:"childHealthchecks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#child_health_threshold TfHealthCheck#child_health_threshold}.
	// Experimental.
	ChildHealthThreshold *float64 `field:"optional" json:"childHealthThreshold" yaml:"childHealthThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#cloudwatch_alarm_name TfHealthCheck#cloudwatch_alarm_name}.
	// Experimental.
	CloudwatchAlarmName *string `field:"optional" json:"cloudwatchAlarmName" yaml:"cloudwatchAlarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#cloudwatch_alarm_region TfHealthCheck#cloudwatch_alarm_region}.
	// Experimental.
	CloudwatchAlarmRegion *string `field:"optional" json:"cloudwatchAlarmRegion" yaml:"cloudwatchAlarmRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#disabled TfHealthCheck#disabled}.
	// Experimental.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#enable_sni TfHealthCheck#enable_sni}.
	// Experimental.
	EnableSni interface{} `field:"optional" json:"enableSni" yaml:"enableSni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#failure_threshold TfHealthCheck#failure_threshold}.
	// Experimental.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#fqdn TfHealthCheck#fqdn}.
	// Experimental.
	Fqdn *string `field:"optional" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#id TfHealthCheck#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#insufficient_data_health_status TfHealthCheck#insufficient_data_health_status}.
	// Experimental.
	InsufficientDataHealthStatus *string `field:"optional" json:"insufficientDataHealthStatus" yaml:"insufficientDataHealthStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#invert_healthcheck TfHealthCheck#invert_healthcheck}.
	// Experimental.
	InvertHealthcheck interface{} `field:"optional" json:"invertHealthcheck" yaml:"invertHealthcheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#ip_address TfHealthCheck#ip_address}.
	// Experimental.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#measure_latency TfHealthCheck#measure_latency}.
	// Experimental.
	MeasureLatency interface{} `field:"optional" json:"measureLatency" yaml:"measureLatency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#port TfHealthCheck#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#reference_name TfHealthCheck#reference_name}.
	// Experimental.
	ReferenceName *string `field:"optional" json:"referenceName" yaml:"referenceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#regions TfHealthCheck#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#request_interval TfHealthCheck#request_interval}.
	// Experimental.
	RequestInterval *float64 `field:"optional" json:"requestInterval" yaml:"requestInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#resource_path TfHealthCheck#resource_path}.
	// Experimental.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#routing_control_arn TfHealthCheck#routing_control_arn}.
	// Experimental.
	RoutingControlArn *string `field:"optional" json:"routingControlArn" yaml:"routingControlArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#search_string TfHealthCheck#search_string}.
	// Experimental.
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#tags TfHealthCheck#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#tags_all TfHealthCheck#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check#triggers TfHealthCheck#triggers}.
	// Experimental.
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

