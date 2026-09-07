package mskconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/mskconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/mskconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnector_AutoscalingPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConnector_AutoscalingProperty
	// Experimental.
	SetInternalValue(val *AwsConnector_AutoscalingProperty)
	// Experimental.
	MaxWorkerCount() *float64
	// Experimental.
	SetMaxWorkerCount(val *float64)
	// Experimental.
	MaxWorkerCountInput() *float64
	// Experimental.
	McuCount() *float64
	// Experimental.
	SetMcuCount(val *float64)
	// Experimental.
	McuCountInput() *float64
	// Experimental.
	MinWorkerCount() *float64
	// Experimental.
	SetMinWorkerCount(val *float64)
	// Experimental.
	MinWorkerCountInput() *float64
	// Experimental.
	ScaleInPolicy() AwsConnector_ScaleInPolicyPropertyOutputReference
	// Experimental.
	ScaleInPolicyInput() *AwsConnector_ScaleInPolicyProperty
	// Experimental.
	ScaleOutPolicy() AwsConnector_ScaleOutPolicyPropertyOutputReference
	// Experimental.
	ScaleOutPolicyInput() *AwsConnector_ScaleOutPolicyProperty
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
	PutScaleInPolicy(value *AwsConnector_ScaleInPolicyProperty)
	// Experimental.
	PutScaleOutPolicy(value *AwsConnector_ScaleOutPolicyProperty)
	// Experimental.
	ResetMcuCount()
	// Experimental.
	ResetScaleInPolicy()
	// Experimental.
	ResetScaleOutPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConnector_AutoscalingPropertyOutputReference
type jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) InternalValue() *AwsConnector_AutoscalingProperty {
	var returns *AwsConnector_AutoscalingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) MaxWorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) MaxWorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) McuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mcuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) McuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mcuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) MinWorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) MinWorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ScaleInPolicy() AwsConnector_ScaleInPolicyPropertyOutputReference {
	var returns AwsConnector_ScaleInPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"scaleInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ScaleInPolicyInput() *AwsConnector_ScaleInPolicyProperty {
	var returns *AwsConnector_ScaleInPolicyProperty
	_jsii_.Get(
		j,
		"scaleInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ScaleOutPolicy() AwsConnector_ScaleOutPolicyPropertyOutputReference {
	var returns AwsConnector_ScaleOutPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"scaleOutPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ScaleOutPolicyInput() *AwsConnector_ScaleOutPolicyProperty {
	var returns *AwsConnector_ScaleOutPolicyProperty
	_jsii_.Get(
		j,
		"scaleOutPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnector_AutoscalingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnector_AutoscalingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnector_AutoscalingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-mskconnect.AwsConnector.AutoscalingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnector_AutoscalingPropertyOutputReference_Override(a AwsConnector_AutoscalingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mskconnect.AwsConnector.AutoscalingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetInternalValue(val *AwsConnector_AutoscalingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetMaxWorkerCount(val *float64) {
	if err := j.validateSetMaxWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkerCount",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetMcuCount(val *float64) {
	if err := j.validateSetMcuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mcuCount",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetMinWorkerCount(val *float64) {
	if err := j.validateSetMinWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWorkerCount",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) PutScaleInPolicy(value *AwsConnector_ScaleInPolicyProperty) {
	if err := a.validatePutScaleInPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScaleInPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) PutScaleOutPolicy(value *AwsConnector_ScaleOutPolicyProperty) {
	if err := a.validatePutScaleOutPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScaleOutPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ResetMcuCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMcuCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ResetScaleInPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleInPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ResetScaleOutPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleOutPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnector_AutoscalingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

