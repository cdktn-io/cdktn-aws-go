package awselbclassic

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfElbConfig struct {
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
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#listener TfElb#listener}
	// Experimental.
	Listener interface{} `field:"required" json:"listener" yaml:"listener"`
	// access_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#access_logs TfElb#access_logs}
	// Experimental.
	AccessLogs *TfElb_AccessLogsProperty `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#availability_zones TfElb#availability_zones}.
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#connection_draining TfElb#connection_draining}.
	// Experimental.
	ConnectionDraining interface{} `field:"optional" json:"connectionDraining" yaml:"connectionDraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#connection_draining_timeout TfElb#connection_draining_timeout}.
	// Experimental.
	ConnectionDrainingTimeout *float64 `field:"optional" json:"connectionDrainingTimeout" yaml:"connectionDrainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#cross_zone_load_balancing TfElb#cross_zone_load_balancing}.
	// Experimental.
	CrossZoneLoadBalancing interface{} `field:"optional" json:"crossZoneLoadBalancing" yaml:"crossZoneLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#desync_mitigation_mode TfElb#desync_mitigation_mode}.
	// Experimental.
	DesyncMitigationMode *string `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#health_check TfElb#health_check}
	// Experimental.
	HealthCheck *TfElb_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#id TfElb#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#idle_timeout TfElb#idle_timeout}.
	// Experimental.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#instances TfElb#instances}.
	// Experimental.
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#internal TfElb#internal}.
	// Experimental.
	Internal interface{} `field:"optional" json:"internal" yaml:"internal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#name TfElb#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#name_prefix TfElb#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#region TfElb#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#security_groups TfElb#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#source_security_group TfElb#source_security_group}.
	// Experimental.
	SourceSecurityGroup *string `field:"optional" json:"sourceSecurityGroup" yaml:"sourceSecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#subnets TfElb#subnets}.
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#tags TfElb#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#tags_all TfElb#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#timeouts TfElb#timeouts}
	// Experimental.
	Timeouts *TfElb_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

