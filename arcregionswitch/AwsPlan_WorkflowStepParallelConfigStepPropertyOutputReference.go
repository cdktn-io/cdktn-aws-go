package arcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/arcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/arcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArcRoutingControlConfig() AwsPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigPropertyList
	// Experimental.
	ArcRoutingControlConfigInput() interface{}
	// Experimental.
	AuroraProvisionedScalingConfig() AwsPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyList
	// Experimental.
	AuroraProvisionedScalingConfigInput() interface{}
	// Experimental.
	AuroraServerlessScalingConfig() AwsPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigPropertyList
	// Experimental.
	AuroraServerlessScalingConfigInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomActionLambdaConfig() AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyList
	// Experimental.
	CustomActionLambdaConfigInput() interface{}
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DocumentDbConfig() AwsPlan_WorkflowStepParallelConfigStepDocumentDbConfigPropertyList
	// Experimental.
	DocumentDbConfigInput() interface{}
	// Experimental.
	Ec2AsgCapacityIncreaseConfig() AwsPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyList
	// Experimental.
	Ec2AsgCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EcsCapacityIncreaseConfig() AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyList
	// Experimental.
	EcsCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EksResourceScalingConfig() AwsPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigPropertyList
	// Experimental.
	EksResourceScalingConfigInput() interface{}
	// Experimental.
	ExecutionApprovalConfig() AwsPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigPropertyList
	// Experimental.
	ExecutionApprovalConfigInput() interface{}
	// Experimental.
	ExecutionBlockType() *string
	// Experimental.
	SetExecutionBlockType(val *string)
	// Experimental.
	ExecutionBlockTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GlobalAuroraConfig() AwsPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigPropertyList
	// Experimental.
	GlobalAuroraConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaEventSourceMappingConfig() AwsPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigPropertyList
	// Experimental.
	LambdaEventSourceMappingConfigInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NeptuneGlobalDatabaseConfig() AwsPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigPropertyList
	// Experimental.
	NeptuneGlobalDatabaseConfigInput() interface{}
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfig() AwsPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfigInput() interface{}
	// Experimental.
	RdsPromoteReadReplicaConfig() AwsPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigPropertyList
	// Experimental.
	RdsPromoteReadReplicaConfigInput() interface{}
	// Experimental.
	RegionSwitchPlanConfig() AwsPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigPropertyList
	// Experimental.
	RegionSwitchPlanConfigInput() interface{}
	// Experimental.
	Route53HealthCheckConfig() AwsPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigPropertyList
	// Experimental.
	Route53HealthCheckConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutArcRoutingControlConfig(value interface{})
	// Experimental.
	PutAuroraProvisionedScalingConfig(value interface{})
	// Experimental.
	PutAuroraServerlessScalingConfig(value interface{})
	// Experimental.
	PutCustomActionLambdaConfig(value interface{})
	// Experimental.
	PutDocumentDbConfig(value interface{})
	// Experimental.
	PutEc2AsgCapacityIncreaseConfig(value interface{})
	// Experimental.
	PutEcsCapacityIncreaseConfig(value interface{})
	// Experimental.
	PutEksResourceScalingConfig(value interface{})
	// Experimental.
	PutExecutionApprovalConfig(value interface{})
	// Experimental.
	PutGlobalAuroraConfig(value interface{})
	// Experimental.
	PutLambdaEventSourceMappingConfig(value interface{})
	// Experimental.
	PutNeptuneGlobalDatabaseConfig(value interface{})
	// Experimental.
	PutRdsCreateCrossRegionReadReplicaConfig(value interface{})
	// Experimental.
	PutRdsPromoteReadReplicaConfig(value interface{})
	// Experimental.
	PutRegionSwitchPlanConfig(value interface{})
	// Experimental.
	PutRoute53HealthCheckConfig(value interface{})
	// Experimental.
	ResetArcRoutingControlConfig()
	// Experimental.
	ResetAuroraProvisionedScalingConfig()
	// Experimental.
	ResetAuroraServerlessScalingConfig()
	// Experimental.
	ResetCustomActionLambdaConfig()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDocumentDbConfig()
	// Experimental.
	ResetEc2AsgCapacityIncreaseConfig()
	// Experimental.
	ResetEcsCapacityIncreaseConfig()
	// Experimental.
	ResetEksResourceScalingConfig()
	// Experimental.
	ResetExecutionApprovalConfig()
	// Experimental.
	ResetGlobalAuroraConfig()
	// Experimental.
	ResetLambdaEventSourceMappingConfig()
	// Experimental.
	ResetNeptuneGlobalDatabaseConfig()
	// Experimental.
	ResetRdsCreateCrossRegionReadReplicaConfig()
	// Experimental.
	ResetRdsPromoteReadReplicaConfig()
	// Experimental.
	ResetRegionSwitchPlanConfig()
	// Experimental.
	ResetRoute53HealthCheckConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference
type jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ArcRoutingControlConfig() AwsPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigPropertyList
	_jsii_.Get(
		j,
		"arcRoutingControlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ArcRoutingControlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arcRoutingControlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraProvisionedScalingConfig() AwsPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraProvisionedScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraServerlessScalingConfig() AwsPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraServerlessScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) CustomActionLambdaConfig() AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyList
	_jsii_.Get(
		j,
		"customActionLambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) CustomActionLambdaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customActionLambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) DocumentDbConfig() AwsPlan_WorkflowStepParallelConfigStepDocumentDbConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepDocumentDbConfigPropertyList
	_jsii_.Get(
		j,
		"documentDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) DocumentDbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentDbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfig() AwsPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EcsCapacityIncreaseConfig() AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EcsCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EksResourceScalingConfig() AwsPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigPropertyList
	_jsii_.Get(
		j,
		"eksResourceScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EksResourceScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksResourceScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionApprovalConfig() AwsPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigPropertyList
	_jsii_.Get(
		j,
		"executionApprovalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionApprovalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionApprovalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionBlockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionBlockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GlobalAuroraConfig() AwsPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigPropertyList
	_jsii_.Get(
		j,
		"globalAuroraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GlobalAuroraConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalAuroraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) LambdaEventSourceMappingConfig() AwsPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigPropertyList
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) LambdaEventSourceMappingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) NeptuneGlobalDatabaseConfig() AwsPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigPropertyList
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) NeptuneGlobalDatabaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfig() AwsPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsPromoteReadReplicaConfig() AwsPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsPromoteReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RegionSwitchPlanConfig() AwsPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigPropertyList
	_jsii_.Get(
		j,
		"regionSwitchPlanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RegionSwitchPlanConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionSwitchPlanConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Route53HealthCheckConfig() AwsPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigPropertyList
	_jsii_.Get(
		j,
		"route53HealthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Route53HealthCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"route53HealthCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPlan_WorkflowStepParallelConfigStepPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepParallelConfigStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference_Override(a AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepParallelConfigStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetExecutionBlockType(val *string) {
	if err := j.validateSetExecutionBlockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionBlockType",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutArcRoutingControlConfig(value interface{}) {
	if err := a.validatePutArcRoutingControlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArcRoutingControlConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutAuroraProvisionedScalingConfig(value interface{}) {
	if err := a.validatePutAuroraProvisionedScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuroraProvisionedScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutAuroraServerlessScalingConfig(value interface{}) {
	if err := a.validatePutAuroraServerlessScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuroraServerlessScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutCustomActionLambdaConfig(value interface{}) {
	if err := a.validatePutCustomActionLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomActionLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutDocumentDbConfig(value interface{}) {
	if err := a.validatePutDocumentDbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDocumentDbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutEc2AsgCapacityIncreaseConfig(value interface{}) {
	if err := a.validatePutEc2AsgCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2AsgCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutEcsCapacityIncreaseConfig(value interface{}) {
	if err := a.validatePutEcsCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutEksResourceScalingConfig(value interface{}) {
	if err := a.validatePutEksResourceScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEksResourceScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutExecutionApprovalConfig(value interface{}) {
	if err := a.validatePutExecutionApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExecutionApprovalConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutGlobalAuroraConfig(value interface{}) {
	if err := a.validatePutGlobalAuroraConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalAuroraConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutLambdaEventSourceMappingConfig(value interface{}) {
	if err := a.validatePutLambdaEventSourceMappingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaEventSourceMappingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutNeptuneGlobalDatabaseConfig(value interface{}) {
	if err := a.validatePutNeptuneGlobalDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNeptuneGlobalDatabaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRdsCreateCrossRegionReadReplicaConfig(value interface{}) {
	if err := a.validatePutRdsCreateCrossRegionReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRdsCreateCrossRegionReadReplicaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRdsPromoteReadReplicaConfig(value interface{}) {
	if err := a.validatePutRdsPromoteReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRdsPromoteReadReplicaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRegionSwitchPlanConfig(value interface{}) {
	if err := a.validatePutRegionSwitchPlanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegionSwitchPlanConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRoute53HealthCheckConfig(value interface{}) {
	if err := a.validatePutRoute53HealthCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoute53HealthCheckConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetArcRoutingControlConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetArcRoutingControlConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetAuroraProvisionedScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAuroraProvisionedScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetAuroraServerlessScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAuroraServerlessScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetCustomActionLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomActionLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetDocumentDbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentDbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetEc2AsgCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2AsgCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetEcsCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetEksResourceScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEksResourceScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetExecutionApprovalConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionApprovalConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetGlobalAuroraConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalAuroraConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetLambdaEventSourceMappingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaEventSourceMappingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetNeptuneGlobalDatabaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNeptuneGlobalDatabaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRdsCreateCrossRegionReadReplicaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsCreateCrossRegionReadReplicaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRdsPromoteReadReplicaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsPromoteReadReplicaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRegionSwitchPlanConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionSwitchPlanConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRoute53HealthCheckConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRoute53HealthCheckConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

