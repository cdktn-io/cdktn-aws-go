package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference interface {
	cdktn.ComplexObject
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
	CrossAccountRole() *string
	// Experimental.
	SetCrossAccountRole(val *string)
	// Experimental.
	CrossAccountRoleInput() *string
	// Experimental.
	ExternalId() *string
	// Experimental.
	SetExternalId(val *string)
	// Experimental.
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GlobalClusterIdentifier() *string
	// Experimental.
	SetGlobalClusterIdentifier(val *string)
	// Experimental.
	GlobalClusterIdentifierInput() *string
	// Experimental.
	InstanceArns() *map[string]*string
	// Experimental.
	SetInstanceArns(val *map[string]*string)
	// Experimental.
	InstanceArnsInput() *map[string]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RegionDatabaseClusterArns() *map[string]*string
	// Experimental.
	SetRegionDatabaseClusterArns(val *map[string]*string)
	// Experimental.
	RegionDatabaseClusterArnsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutMinutes() *float64
	// Experimental.
	SetTimeoutMinutes(val *float64)
	// Experimental.
	TimeoutMinutesInput() *float64
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
	ResetCrossAccountRole()
	// Experimental.
	ResetExternalId()
	// Experimental.
	ResetTimeoutMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference
type jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) CrossAccountRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossAccountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) CrossAccountRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossAccountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) InstanceArns() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"instanceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) InstanceArnsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"instanceArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) RegionDatabaseClusterArns() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"regionDatabaseClusterArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) RegionDatabaseClusterArnsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"regionDatabaseClusterArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsArcregionswitchPlan.WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference_Override(a AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsArcregionswitchPlan.WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetCrossAccountRole(val *string) {
	if err := j.validateSetCrossAccountRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossAccountRole",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetInstanceArns(val *map[string]*string) {
	if err := j.validateSetInstanceArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceArns",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetRegionDatabaseClusterArns(val *map[string]*string) {
	if err := j.validateSetRegionDatabaseClusterArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionDatabaseClusterArns",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ResetCrossAccountRole() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossAccountRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

