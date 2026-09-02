package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArcRoutingControlConfig() TfPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigPropertyList
	// Experimental.
	ArcRoutingControlConfigInput() interface{}
	// Experimental.
	AuroraProvisionedScalingConfig() TfPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyList
	// Experimental.
	AuroraProvisionedScalingConfigInput() interface{}
	// Experimental.
	AuroraServerlessScalingConfig() TfPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigPropertyList
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
	CustomActionLambdaConfig() TfPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyList
	// Experimental.
	CustomActionLambdaConfigInput() interface{}
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DocumentDbConfig() TfPlan_WorkflowStepParallelConfigStepDocumentDbConfigPropertyList
	// Experimental.
	DocumentDbConfigInput() interface{}
	// Experimental.
	Ec2AsgCapacityIncreaseConfig() TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyList
	// Experimental.
	Ec2AsgCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EcsCapacityIncreaseConfig() TfPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyList
	// Experimental.
	EcsCapacityIncreaseConfigInput() interface{}
	// Experimental.
	EksResourceScalingConfig() TfPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigPropertyList
	// Experimental.
	EksResourceScalingConfigInput() interface{}
	// Experimental.
	ExecutionApprovalConfig() TfPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigPropertyList
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
	GlobalAuroraConfig() TfPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigPropertyList
	// Experimental.
	GlobalAuroraConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaEventSourceMappingConfig() TfPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigPropertyList
	// Experimental.
	LambdaEventSourceMappingConfigInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NeptuneGlobalDatabaseConfig() TfPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigPropertyList
	// Experimental.
	NeptuneGlobalDatabaseConfigInput() interface{}
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfig() TfPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfigInput() interface{}
	// Experimental.
	RdsPromoteReadReplicaConfig() TfPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigPropertyList
	// Experimental.
	RdsPromoteReadReplicaConfigInput() interface{}
	// Experimental.
	RegionSwitchPlanConfig() TfPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigPropertyList
	// Experimental.
	RegionSwitchPlanConfigInput() interface{}
	// Experimental.
	Route53HealthCheckConfig() TfPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigPropertyList
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

// The jsii proxy struct for TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference
type jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ArcRoutingControlConfig() TfPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigPropertyList
	_jsii_.Get(
		j,
		"arcRoutingControlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ArcRoutingControlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arcRoutingControlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraProvisionedScalingConfig() TfPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraProvisionedScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraProvisionedScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraServerlessScalingConfig() TfPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigPropertyList
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) AuroraServerlessScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auroraServerlessScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) CustomActionLambdaConfig() TfPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyList
	_jsii_.Get(
		j,
		"customActionLambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) CustomActionLambdaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customActionLambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) DocumentDbConfig() TfPlan_WorkflowStepParallelConfigStepDocumentDbConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepDocumentDbConfigPropertyList
	_jsii_.Get(
		j,
		"documentDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) DocumentDbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentDbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfig() TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Ec2AsgCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EcsCapacityIncreaseConfig() TfPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyList
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EcsCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EksResourceScalingConfig() TfPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigPropertyList
	_jsii_.Get(
		j,
		"eksResourceScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) EksResourceScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksResourceScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionApprovalConfig() TfPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigPropertyList
	_jsii_.Get(
		j,
		"executionApprovalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionApprovalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionApprovalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionBlockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ExecutionBlockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GlobalAuroraConfig() TfPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigPropertyList
	_jsii_.Get(
		j,
		"globalAuroraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GlobalAuroraConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalAuroraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) LambdaEventSourceMappingConfig() TfPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigPropertyList
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) LambdaEventSourceMappingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaEventSourceMappingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) NeptuneGlobalDatabaseConfig() TfPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigPropertyList
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) NeptuneGlobalDatabaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneGlobalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfig() TfPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsCreateCrossRegionReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsCreateCrossRegionReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsPromoteReadReplicaConfig() TfPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigPropertyList
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RdsPromoteReadReplicaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsPromoteReadReplicaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RegionSwitchPlanConfig() TfPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigPropertyList
	_jsii_.Get(
		j,
		"regionSwitchPlanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) RegionSwitchPlanConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionSwitchPlanConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Route53HealthCheckConfig() TfPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigPropertyList
	_jsii_.Get(
		j,
		"route53HealthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Route53HealthCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"route53HealthCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPlan_WorkflowStepParallelConfigStepPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPlan_WorkflowStepParallelConfigStepPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepParallelConfigStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPlan_WorkflowStepParallelConfigStepPropertyOutputReference_Override(t TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepParallelConfigStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetExecutionBlockType(val *string) {
	if err := j.validateSetExecutionBlockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionBlockType",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutArcRoutingControlConfig(value interface{}) {
	if err := t.validatePutArcRoutingControlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArcRoutingControlConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutAuroraProvisionedScalingConfig(value interface{}) {
	if err := t.validatePutAuroraProvisionedScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuroraProvisionedScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutAuroraServerlessScalingConfig(value interface{}) {
	if err := t.validatePutAuroraServerlessScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuroraServerlessScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutCustomActionLambdaConfig(value interface{}) {
	if err := t.validatePutCustomActionLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomActionLambdaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutDocumentDbConfig(value interface{}) {
	if err := t.validatePutDocumentDbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDocumentDbConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutEc2AsgCapacityIncreaseConfig(value interface{}) {
	if err := t.validatePutEc2AsgCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2AsgCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutEcsCapacityIncreaseConfig(value interface{}) {
	if err := t.validatePutEcsCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcsCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutEksResourceScalingConfig(value interface{}) {
	if err := t.validatePutEksResourceScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEksResourceScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutExecutionApprovalConfig(value interface{}) {
	if err := t.validatePutExecutionApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExecutionApprovalConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutGlobalAuroraConfig(value interface{}) {
	if err := t.validatePutGlobalAuroraConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGlobalAuroraConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutLambdaEventSourceMappingConfig(value interface{}) {
	if err := t.validatePutLambdaEventSourceMappingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaEventSourceMappingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutNeptuneGlobalDatabaseConfig(value interface{}) {
	if err := t.validatePutNeptuneGlobalDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNeptuneGlobalDatabaseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRdsCreateCrossRegionReadReplicaConfig(value interface{}) {
	if err := t.validatePutRdsCreateCrossRegionReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRdsCreateCrossRegionReadReplicaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRdsPromoteReadReplicaConfig(value interface{}) {
	if err := t.validatePutRdsPromoteReadReplicaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRdsPromoteReadReplicaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRegionSwitchPlanConfig(value interface{}) {
	if err := t.validatePutRegionSwitchPlanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegionSwitchPlanConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) PutRoute53HealthCheckConfig(value interface{}) {
	if err := t.validatePutRoute53HealthCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRoute53HealthCheckConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetArcRoutingControlConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetArcRoutingControlConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetAuroraProvisionedScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAuroraProvisionedScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetAuroraServerlessScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAuroraServerlessScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetCustomActionLambdaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomActionLambdaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetDocumentDbConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentDbConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetEc2AsgCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2AsgCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetEcsCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEcsCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetEksResourceScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEksResourceScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetExecutionApprovalConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionApprovalConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetGlobalAuroraConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetGlobalAuroraConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetLambdaEventSourceMappingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaEventSourceMappingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetNeptuneGlobalDatabaseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetNeptuneGlobalDatabaseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRdsCreateCrossRegionReadReplicaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsCreateCrossRegionReadReplicaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRdsPromoteReadReplicaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsPromoteReadReplicaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRegionSwitchPlanConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRegionSwitchPlanConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ResetRoute53HealthCheckConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRoute53HealthCheckConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

