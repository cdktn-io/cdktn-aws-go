package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPlan_WorkflowStepPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArcRoutingControlConfig() TfPlan_WorkflowStepArcRoutingControlConfigPropertyList
	// Experimental.
	ArcRoutingControlConfigInput() interface{}
	// Experimental.
	AuroraProvisionedScalingConfig() TfPlan_WorkflowStepAuroraProvisionedScalingConfigPropertyList
	// Experimental.
	AuroraProvisionedScalingConfigInput() interface{}
	// Experimental.
	AuroraServerlessScalingConfig() TfPlan_WorkflowStepAuroraServerlessScalingConfigPropertyList
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
	CustomActionLambdaConfig() TfPlan_WorkflowStepCustomActionLambdaConfigPropertyList
	// Experimental.
	CustomActionLambdaConfigInput() interface{}
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DocumentDbConfig() TfPlan_WorkflowStepDocumentDbConfigPropertyList
	// Experimental.
	DocumentDbConfigInput() interface{}
	// Experimental.
	Ec2AsgCapacityIncreaseConfig() TfPlan_WorkflowStepEc2AsgCapacityIncreaseConfigPropertyList
	// Experimental.
	Ec2AsgCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EcsCapacityIncreaseConfig() TfPlan_WorkflowStepEcsCapacityIncreaseConfigPropertyList
	// Experimental.
	EcsCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EksResourceScalingConfig() TfPlan_WorkflowStepEksResourceScalingConfigPropertyList
	// Experimental.
	EksResourceScalingConfigInput() interface{}
	// Experimental.
	ExecutionApprovalConfig() TfPlan_WorkflowStepExecutionApprovalConfigPropertyList
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
	GlobalAuroraConfig() TfPlan_WorkflowStepGlobalAuroraConfigPropertyList
	// Experimental.
	GlobalAuroraConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaEventSourceMappingConfig() TfPlan_WorkflowStepLambdaEventSourceMappingConfigPropertyList
	// Experimental.
	LambdaEventSourceMappingConfigInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NeptuneGlobalDatabaseConfig() TfPlan_WorkflowStepNeptuneGlobalDatabaseConfigPropertyList
	// Experimental.
	NeptuneGlobalDatabaseConfigInput() interface{}
	// Experimental.
	ParallelConfig() TfPlan_ParallelConfigPropertyList
	// Experimental.
	ParallelConfigInput() interface{}
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfig() TfPlan_WorkflowStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfigInput() interface{}
	// Experimental.
	RdsPromoteReadReplicaConfig() TfPlan_WorkflowStepRdsPromoteReadReplicaConfigPropertyList
	// Experimental.
	RdsPromoteReadReplicaConfigInput() interface{}
	// Experimental.
	RegionSwitchPlanConfig() TfPlan_WorkflowStepRegionSwitchPlanConfigPropertyList
	// Experimental.
	RegionSwitchPlanConfigInput() interface{}
	// Experimental.
	Route53HealthCheckConfig() TfPlan_WorkflowStepRoute53HealthCheckConfigPropertyList
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

// The jsii proxy struct for TfPlan_WorkflowStepPropertyOutputReference
type jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ArcRoutingControlConfig() TfPlan_WorkflowStepArcRoutingControlConfigPropertyList {
	var returns TfPlan_WorkflowStepArcRoutingControlConfigPropertyList
	_jsii_.Get(
		j,
		"arcRoutingControlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ArcRoutingControlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arcRoutingControlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) AuroraProvisionedScalingConfig() TfPlan_WorkflowStepAuroraProvisionedScalingConfigPropertyList {
	var returns TfPlan_WorkflowStepAuroraProvisionedScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) AuroraProvisionedScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) AuroraServerlessScalingConfig() TfPlan_WorkflowStepAuroraServerlessScalingConfigPropertyList {
	var returns TfPlan_WorkflowStepAuroraServerlessScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) AuroraServerlessScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) CustomActionLambdaConfig() TfPlan_WorkflowStepCustomActionLambdaConfigPropertyList {
	var returns TfPlan_WorkflowStepCustomActionLambdaConfigPropertyList
	_jsii_.Get(
		j,
		"customActionLambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) CustomActionLambdaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customActionLambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) DocumentDbConfig() TfPlan_WorkflowStepDocumentDbConfigPropertyList {
	var returns TfPlan_WorkflowStepDocumentDbConfigPropertyList
	_jsii_.Get(
		j,
		"documentDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) DocumentDbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentDbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfig() TfPlan_WorkflowStepEc2AsgCapacityIncreaseConfigPropertyList {
	var returns TfPlan_WorkflowStepEc2AsgCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) EcsCapacityIncreaseConfig() TfPlan_WorkflowStepEcsCapacityIncreaseConfigPropertyList {
	var returns TfPlan_WorkflowStepEcsCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) EcsCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) EksResourceScalingConfig() TfPlan_WorkflowStepEksResourceScalingConfigPropertyList {
	var returns TfPlan_WorkflowStepEksResourceScalingConfigPropertyList
	_jsii_.Get(
		j,
		"eksResourceScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) EksResourceScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksResourceScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ExecutionApprovalConfig() TfPlan_WorkflowStepExecutionApprovalConfigPropertyList {
	var returns TfPlan_WorkflowStepExecutionApprovalConfigPropertyList
	_jsii_.Get(
		j,
		"executionApprovalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ExecutionApprovalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionApprovalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ExecutionBlockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ExecutionBlockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GlobalAuroraConfig() TfPlan_WorkflowStepGlobalAuroraConfigPropertyList {
	var returns TfPlan_WorkflowStepGlobalAuroraConfigPropertyList
	_jsii_.Get(
		j,
		"globalAuroraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GlobalAuroraConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalAuroraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) LambdaEventSourceMappingConfig() TfPlan_WorkflowStepLambdaEventSourceMappingConfigPropertyList {
	var returns TfPlan_WorkflowStepLambdaEventSourceMappingConfigPropertyList
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) LambdaEventSourceMappingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) NeptuneGlobalDatabaseConfig() TfPlan_WorkflowStepNeptuneGlobalDatabaseConfigPropertyList {
	var returns TfPlan_WorkflowStepNeptuneGlobalDatabaseConfigPropertyList
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) NeptuneGlobalDatabaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ParallelConfig() TfPlan_ParallelConfigPropertyList {
	var returns TfPlan_ParallelConfigPropertyList
	_jsii_.Get(
		j,
		"parallelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ParallelConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parallelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfig() TfPlan_WorkflowStepRdsCreateCrossRegionReadReplicaConfigPropertyList {
	var returns TfPlan_WorkflowStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) RdsPromoteReadReplicaConfig() TfPlan_WorkflowStepRdsPromoteReadReplicaConfigPropertyList {
	var returns TfPlan_WorkflowStepRdsPromoteReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) RdsPromoteReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) RegionSwitchPlanConfig() TfPlan_WorkflowStepRegionSwitchPlanConfigPropertyList {
	var returns TfPlan_WorkflowStepRegionSwitchPlanConfigPropertyList
	_jsii_.Get(
		j,
		"regionSwitchPlanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) RegionSwitchPlanConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionSwitchPlanConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Route53HealthCheckConfig() TfPlan_WorkflowStepRoute53HealthCheckConfigPropertyList {
	var returns TfPlan_WorkflowStepRoute53HealthCheckConfigPropertyList
	_jsii_.Get(
		j,
		"route53HealthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Route53HealthCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"route53HealthCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPlan_WorkflowStepPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPlan_WorkflowStepPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPlan_WorkflowStepPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPlan_WorkflowStepPropertyOutputReference_Override(t TfPlan_WorkflowStepPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetExecutionBlockType(val *string) {
	if err := j.validateSetExecutionBlockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionBlockType",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutArcRoutingControlConfig(value interface{}) {
	if err := t.validatePutArcRoutingControlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArcRoutingControlConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutAuroraProvisionedScalingConfig(value interface{}) {
	if err := t.validatePutAuroraProvisionedScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuroraProvisionedScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutAuroraServerlessScalingConfig(value interface{}) {
	if err := t.validatePutAuroraServerlessScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuroraServerlessScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutCustomActionLambdaConfig(value interface{}) {
	if err := t.validatePutCustomActionLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomActionLambdaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutDocumentDbConfig(value interface{}) {
	if err := t.validatePutDocumentDbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDocumentDbConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutEc2AsgCapacityIncreaseConfig(value interface{}) {
	if err := t.validatePutEc2AsgCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2AsgCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutEcsCapacityIncreaseConfig(value interface{}) {
	if err := t.validatePutEcsCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcsCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutEksResourceScalingConfig(value interface{}) {
	if err := t.validatePutEksResourceScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEksResourceScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutExecutionApprovalConfig(value interface{}) {
	if err := t.validatePutExecutionApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExecutionApprovalConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutGlobalAuroraConfig(value interface{}) {
	if err := t.validatePutGlobalAuroraConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGlobalAuroraConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutLambdaEventSourceMappingConfig(value interface{}) {
	if err := t.validatePutLambdaEventSourceMappingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaEventSourceMappingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutNeptuneGlobalDatabaseConfig(value interface{}) {
	if err := t.validatePutNeptuneGlobalDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNeptuneGlobalDatabaseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutParallelConfig(value interface{}) {
	if err := t.validatePutParallelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParallelConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutRdsCreateCrossRegionReadReplicaConfig(value interface{}) {
	if err := t.validatePutRdsCreateCrossRegionReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRdsCreateCrossRegionReadReplicaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutRdsPromoteReadReplicaConfig(value interface{}) {
	if err := t.validatePutRdsPromoteReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRdsPromoteReadReplicaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutRegionSwitchPlanConfig(value interface{}) {
	if err := t.validatePutRegionSwitchPlanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegionSwitchPlanConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) PutRoute53HealthCheckConfig(value interface{}) {
	if err := t.validatePutRoute53HealthCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRoute53HealthCheckConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetArcRoutingControlConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetArcRoutingControlConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetAuroraProvisionedScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAuroraProvisionedScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetAuroraServerlessScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAuroraServerlessScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetCustomActionLambdaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomActionLambdaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetDocumentDbConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentDbConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetEc2AsgCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2AsgCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetEcsCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEcsCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetEksResourceScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEksResourceScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetExecutionApprovalConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionApprovalConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetGlobalAuroraConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetGlobalAuroraConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetLambdaEventSourceMappingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaEventSourceMappingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetNeptuneGlobalDatabaseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetNeptuneGlobalDatabaseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetParallelConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetRdsCreateCrossRegionReadReplicaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsCreateCrossRegionReadReplicaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetRdsPromoteReadReplicaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsPromoteReadReplicaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetRegionSwitchPlanConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRegionSwitchPlanConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ResetRoute53HealthCheckConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRoute53HealthCheckConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

