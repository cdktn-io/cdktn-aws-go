package arcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/arcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/arcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPlan_WorkflowStepPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArcRoutingControlConfig() AwsPlan_WorkflowStepArcRoutingControlConfigPropertyList
	// Experimental.
	ArcRoutingControlConfigInput() interface{}
	// Experimental.
	AuroraProvisionedScalingConfig() AwsPlan_WorkflowStepAuroraProvisionedScalingConfigPropertyList
	// Experimental.
	AuroraProvisionedScalingConfigInput() interface{}
	// Experimental.
	AuroraServerlessScalingConfig() AwsPlan_WorkflowStepAuroraServerlessScalingConfigPropertyList
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
	CustomActionLambdaConfig() AwsPlan_WorkflowStepCustomActionLambdaConfigPropertyList
	// Experimental.
	CustomActionLambdaConfigInput() interface{}
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DocumentDbConfig() AwsPlan_WorkflowStepDocumentDbConfigPropertyList
	// Experimental.
	DocumentDbConfigInput() interface{}
	// Experimental.
	Ec2AsgCapacityIncreaseConfig() AwsPlan_WorkflowStepEc2AsgCapacityIncreaseConfigPropertyList
	// Experimental.
	Ec2AsgCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EcsCapacityIncreaseConfig() AwsPlan_WorkflowStepEcsCapacityIncreaseConfigPropertyList
	// Experimental.
	EcsCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EksResourceScalingConfig() AwsPlan_WorkflowStepEksResourceScalingConfigPropertyList
	// Experimental.
	EksResourceScalingConfigInput() interface{}
	// Experimental.
	ExecutionApprovalConfig() AwsPlan_WorkflowStepExecutionApprovalConfigPropertyList
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
	GlobalAuroraConfig() AwsPlan_WorkflowStepGlobalAuroraConfigPropertyList
	// Experimental.
	GlobalAuroraConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaEventSourceMappingConfig() AwsPlan_WorkflowStepLambdaEventSourceMappingConfigPropertyList
	// Experimental.
	LambdaEventSourceMappingConfigInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NeptuneGlobalDatabaseConfig() AwsPlan_WorkflowStepNeptuneGlobalDatabaseConfigPropertyList
	// Experimental.
	NeptuneGlobalDatabaseConfigInput() interface{}
	// Experimental.
	ParallelConfig() AwsPlan_ParallelConfigPropertyList
	// Experimental.
	ParallelConfigInput() interface{}
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfig() AwsPlan_WorkflowStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfigInput() interface{}
	// Experimental.
	RdsPromoteReadReplicaConfig() AwsPlan_WorkflowStepRdsPromoteReadReplicaConfigPropertyList
	// Experimental.
	RdsPromoteReadReplicaConfigInput() interface{}
	// Experimental.
	RegionSwitchPlanConfig() AwsPlan_WorkflowStepRegionSwitchPlanConfigPropertyList
	// Experimental.
	RegionSwitchPlanConfigInput() interface{}
	// Experimental.
	Route53HealthCheckConfig() AwsPlan_WorkflowStepRoute53HealthCheckConfigPropertyList
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
	PutParallelConfig(value interface{})
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
	ResetParallelConfig()
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

// The jsii proxy struct for AwsPlan_WorkflowStepPropertyOutputReference
type jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ArcRoutingControlConfig() AwsPlan_WorkflowStepArcRoutingControlConfigPropertyList {
	var returns AwsPlan_WorkflowStepArcRoutingControlConfigPropertyList
	_jsii_.Get(
		j,
		"arcRoutingControlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ArcRoutingControlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arcRoutingControlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) AuroraProvisionedScalingConfig() AwsPlan_WorkflowStepAuroraProvisionedScalingConfigPropertyList {
	var returns AwsPlan_WorkflowStepAuroraProvisionedScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) AuroraProvisionedScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) AuroraServerlessScalingConfig() AwsPlan_WorkflowStepAuroraServerlessScalingConfigPropertyList {
	var returns AwsPlan_WorkflowStepAuroraServerlessScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) AuroraServerlessScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) CustomActionLambdaConfig() AwsPlan_WorkflowStepCustomActionLambdaConfigPropertyList {
	var returns AwsPlan_WorkflowStepCustomActionLambdaConfigPropertyList
	_jsii_.Get(
		j,
		"customActionLambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) CustomActionLambdaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customActionLambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) DocumentDbConfig() AwsPlan_WorkflowStepDocumentDbConfigPropertyList {
	var returns AwsPlan_WorkflowStepDocumentDbConfigPropertyList
	_jsii_.Get(
		j,
		"documentDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) DocumentDbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentDbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfig() AwsPlan_WorkflowStepEc2AsgCapacityIncreaseConfigPropertyList {
	var returns AwsPlan_WorkflowStepEc2AsgCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) EcsCapacityIncreaseConfig() AwsPlan_WorkflowStepEcsCapacityIncreaseConfigPropertyList {
	var returns AwsPlan_WorkflowStepEcsCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) EcsCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) EksResourceScalingConfig() AwsPlan_WorkflowStepEksResourceScalingConfigPropertyList {
	var returns AwsPlan_WorkflowStepEksResourceScalingConfigPropertyList
	_jsii_.Get(
		j,
		"eksResourceScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) EksResourceScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksResourceScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ExecutionApprovalConfig() AwsPlan_WorkflowStepExecutionApprovalConfigPropertyList {
	var returns AwsPlan_WorkflowStepExecutionApprovalConfigPropertyList
	_jsii_.Get(
		j,
		"executionApprovalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ExecutionApprovalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionApprovalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ExecutionBlockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ExecutionBlockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GlobalAuroraConfig() AwsPlan_WorkflowStepGlobalAuroraConfigPropertyList {
	var returns AwsPlan_WorkflowStepGlobalAuroraConfigPropertyList
	_jsii_.Get(
		j,
		"globalAuroraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GlobalAuroraConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalAuroraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) LambdaEventSourceMappingConfig() AwsPlan_WorkflowStepLambdaEventSourceMappingConfigPropertyList {
	var returns AwsPlan_WorkflowStepLambdaEventSourceMappingConfigPropertyList
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) LambdaEventSourceMappingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) NeptuneGlobalDatabaseConfig() AwsPlan_WorkflowStepNeptuneGlobalDatabaseConfigPropertyList {
	var returns AwsPlan_WorkflowStepNeptuneGlobalDatabaseConfigPropertyList
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) NeptuneGlobalDatabaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ParallelConfig() AwsPlan_ParallelConfigPropertyList {
	var returns AwsPlan_ParallelConfigPropertyList
	_jsii_.Get(
		j,
		"parallelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ParallelConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parallelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfig() AwsPlan_WorkflowStepRdsCreateCrossRegionReadReplicaConfigPropertyList {
	var returns AwsPlan_WorkflowStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) RdsPromoteReadReplicaConfig() AwsPlan_WorkflowStepRdsPromoteReadReplicaConfigPropertyList {
	var returns AwsPlan_WorkflowStepRdsPromoteReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) RdsPromoteReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) RegionSwitchPlanConfig() AwsPlan_WorkflowStepRegionSwitchPlanConfigPropertyList {
	var returns AwsPlan_WorkflowStepRegionSwitchPlanConfigPropertyList
	_jsii_.Get(
		j,
		"regionSwitchPlanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) RegionSwitchPlanConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionSwitchPlanConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Route53HealthCheckConfig() AwsPlan_WorkflowStepRoute53HealthCheckConfigPropertyList {
	var returns AwsPlan_WorkflowStepRoute53HealthCheckConfigPropertyList
	_jsii_.Get(
		j,
		"route53HealthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Route53HealthCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"route53HealthCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPlan_WorkflowStepPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPlan_WorkflowStepPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPlan_WorkflowStepPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPlan_WorkflowStepPropertyOutputReference_Override(a AwsPlan_WorkflowStepPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetExecutionBlockType(val *string) {
	if err := j.validateSetExecutionBlockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionBlockType",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutArcRoutingControlConfig(value interface{}) {
	if err := a.validatePutArcRoutingControlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArcRoutingControlConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutAuroraProvisionedScalingConfig(value interface{}) {
	if err := a.validatePutAuroraProvisionedScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuroraProvisionedScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutAuroraServerlessScalingConfig(value interface{}) {
	if err := a.validatePutAuroraServerlessScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuroraServerlessScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutCustomActionLambdaConfig(value interface{}) {
	if err := a.validatePutCustomActionLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomActionLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutDocumentDbConfig(value interface{}) {
	if err := a.validatePutDocumentDbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDocumentDbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutEc2AsgCapacityIncreaseConfig(value interface{}) {
	if err := a.validatePutEc2AsgCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2AsgCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutEcsCapacityIncreaseConfig(value interface{}) {
	if err := a.validatePutEcsCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutEksResourceScalingConfig(value interface{}) {
	if err := a.validatePutEksResourceScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEksResourceScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutExecutionApprovalConfig(value interface{}) {
	if err := a.validatePutExecutionApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExecutionApprovalConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutGlobalAuroraConfig(value interface{}) {
	if err := a.validatePutGlobalAuroraConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalAuroraConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutLambdaEventSourceMappingConfig(value interface{}) {
	if err := a.validatePutLambdaEventSourceMappingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaEventSourceMappingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutNeptuneGlobalDatabaseConfig(value interface{}) {
	if err := a.validatePutNeptuneGlobalDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNeptuneGlobalDatabaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutParallelConfig(value interface{}) {
	if err := a.validatePutParallelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutRdsCreateCrossRegionReadReplicaConfig(value interface{}) {
	if err := a.validatePutRdsCreateCrossRegionReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRdsCreateCrossRegionReadReplicaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutRdsPromoteReadReplicaConfig(value interface{}) {
	if err := a.validatePutRdsPromoteReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRdsPromoteReadReplicaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutRegionSwitchPlanConfig(value interface{}) {
	if err := a.validatePutRegionSwitchPlanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegionSwitchPlanConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) PutRoute53HealthCheckConfig(value interface{}) {
	if err := a.validatePutRoute53HealthCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoute53HealthCheckConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetArcRoutingControlConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetArcRoutingControlConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetAuroraProvisionedScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAuroraProvisionedScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetAuroraServerlessScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAuroraServerlessScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetCustomActionLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomActionLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetDocumentDbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentDbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetEc2AsgCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2AsgCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetEcsCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetEksResourceScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEksResourceScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetExecutionApprovalConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionApprovalConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetGlobalAuroraConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalAuroraConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetLambdaEventSourceMappingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaEventSourceMappingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetNeptuneGlobalDatabaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNeptuneGlobalDatabaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetParallelConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetRdsCreateCrossRegionReadReplicaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsCreateCrossRegionReadReplicaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetRdsPromoteReadReplicaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsPromoteReadReplicaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetRegionSwitchPlanConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionSwitchPlanConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ResetRoute53HealthCheckConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRoute53HealthCheckConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

