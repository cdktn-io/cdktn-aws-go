package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/batch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/batch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobDefinition_EvaluateOnExitPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() *string
	// Experimental.
	SetAction(val *string)
	// Experimental.
	ActionInput() *string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OnExitCode() *string
	// Experimental.
	SetOnExitCode(val *string)
	// Experimental.
	OnExitCodeInput() *string
	// Experimental.
	OnReason() *string
	// Experimental.
	SetOnReason(val *string)
	// Experimental.
	OnReasonInput() *string
	// Experimental.
	OnStatusReason() *string
	// Experimental.
	SetOnStatusReason(val *string)
	// Experimental.
	OnStatusReasonInput() *string
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
	ResetOnExitCode()
	// Experimental.
	ResetOnReason()
	// Experimental.
	ResetOnStatusReason()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsJobDefinition_EvaluateOnExitPropertyOutputReference
type jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) OnExitCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onExitCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) OnExitCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onExitCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) OnReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) OnReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) OnStatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onStatusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) OnStatusReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onStatusReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobDefinition_EvaluateOnExitPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsJobDefinition_EvaluateOnExitPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobDefinition_EvaluateOnExitPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.AwsJobDefinition.EvaluateOnExitPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobDefinition_EvaluateOnExitPropertyOutputReference_Override(a AwsJobDefinition_EvaluateOnExitPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.AwsJobDefinition.EvaluateOnExitPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetOnExitCode(val *string) {
	if err := j.validateSetOnExitCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onExitCode",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetOnReason(val *string) {
	if err := j.validateSetOnReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onReason",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetOnStatusReason(val *string) {
	if err := j.validateSetOnStatusReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStatusReason",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ResetOnExitCode() {
	_jsii_.InvokeVoid(
		a,
		"resetOnExitCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ResetOnReason() {
	_jsii_.InvokeVoid(
		a,
		"resetOnReason",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ResetOnStatusReason() {
	_jsii_.InvokeVoid(
		a,
		"resetOnStatusReason",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobDefinition_EvaluateOnExitPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

