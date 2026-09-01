package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudformation/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference interface {
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
	// Experimental.
	ConcurrencyMode() *string
	// Experimental.
	SetConcurrencyMode(val *string)
	// Experimental.
	ConcurrencyModeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FailureToleranceCount() *float64
	// Experimental.
	SetFailureToleranceCount(val *float64)
	// Experimental.
	FailureToleranceCountInput() *float64
	// Experimental.
	FailureTolerancePercentage() *float64
	// Experimental.
	SetFailureTolerancePercentage(val *float64)
	// Experimental.
	FailureTolerancePercentageInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCloudformationStackInstances_OperationPreferencesProperty
	// Experimental.
	SetInternalValue(val *AwsCloudformationStackInstances_OperationPreferencesProperty)
	// Experimental.
	MaxConcurrentCount() *float64
	// Experimental.
	SetMaxConcurrentCount(val *float64)
	// Experimental.
	MaxConcurrentCountInput() *float64
	// Experimental.
	MaxConcurrentPercentage() *float64
	// Experimental.
	SetMaxConcurrentPercentage(val *float64)
	// Experimental.
	MaxConcurrentPercentageInput() *float64
	// Experimental.
	RegionConcurrencyType() *string
	// Experimental.
	SetRegionConcurrencyType(val *string)
	// Experimental.
	RegionConcurrencyTypeInput() *string
	// Experimental.
	RegionOrder() *[]*string
	// Experimental.
	SetRegionOrder(val *[]*string)
	// Experimental.
	RegionOrderInput() *[]*string
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
	ResetConcurrencyMode()
	// Experimental.
	ResetFailureToleranceCount()
	// Experimental.
	ResetFailureTolerancePercentage()
	// Experimental.
	ResetMaxConcurrentCount()
	// Experimental.
	ResetMaxConcurrentPercentage()
	// Experimental.
	ResetRegionConcurrencyType()
	// Experimental.
	ResetRegionOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference
type jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ConcurrencyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ConcurrencyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) InternalValue() *AwsCloudformationStackInstances_OperationPreferencesProperty {
	var returns *AwsCloudformationStackInstances_OperationPreferencesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) MaxConcurrentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) MaxConcurrentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) MaxConcurrentPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) MaxConcurrentPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) RegionConcurrencyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionConcurrencyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) RegionConcurrencyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionConcurrencyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) RegionOrder() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) RegionOrderInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudformationStackInstances_OperationPreferencesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudformation.AwsCloudformationStackInstances.OperationPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference_Override(a AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudformation.AwsCloudformationStackInstances.OperationPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetConcurrencyMode(val *string) {
	if err := j.validateSetConcurrencyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrencyMode",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetFailureToleranceCount(val *float64) {
	if err := j.validateSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetInternalValue(val *AwsCloudformationStackInstances_OperationPreferencesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetMaxConcurrentCount(val *float64) {
	if err := j.validateSetMaxConcurrentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentCount",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetMaxConcurrentPercentage(val *float64) {
	if err := j.validateSetMaxConcurrentPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetRegionConcurrencyType(val *string) {
	if err := j.validateSetRegionConcurrencyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionConcurrencyType",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetRegionOrder(val *[]*string) {
	if err := j.validateSetRegionOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionOrder",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetConcurrencyMode() {
	_jsii_.InvokeVoid(
		a,
		"resetConcurrencyMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetMaxConcurrentCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetMaxConcurrentPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetRegionConcurrencyType() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionConcurrencyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ResetRegionOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudformationStackInstances_OperationPreferencesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

