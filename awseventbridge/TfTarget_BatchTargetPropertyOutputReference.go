package awseventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridge/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTarget_BatchTargetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArraySize() *float64
	// Experimental.
	SetArraySize(val *float64)
	// Experimental.
	ArraySizeInput() *float64
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
	InternalValue() *TfTarget_BatchTargetProperty
	// Experimental.
	SetInternalValue(val *TfTarget_BatchTargetProperty)
	// Experimental.
	JobAttempts() *float64
	// Experimental.
	SetJobAttempts(val *float64)
	// Experimental.
	JobAttemptsInput() *float64
	// Experimental.
	JobDefinition() *string
	// Experimental.
	SetJobDefinition(val *string)
	// Experimental.
	JobDefinitionInput() *string
	// Experimental.
	JobName() *string
	// Experimental.
	SetJobName(val *string)
	// Experimental.
	JobNameInput() *string
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
	ResetArraySize()
	// Experimental.
	ResetJobAttempts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTarget_BatchTargetPropertyOutputReference
type jsiiProxy_TfTarget_BatchTargetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ArraySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"arraySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ArraySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"arraySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) InternalValue() *TfTarget_BatchTargetProperty {
	var returns *TfTarget_BatchTargetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) JobAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) JobAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) JobDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) JobDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTarget_BatchTargetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTarget_BatchTargetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTarget_BatchTargetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTarget_BatchTargetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfTarget.BatchTargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTarget_BatchTargetPropertyOutputReference_Override(t TfTarget_BatchTargetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfTarget.BatchTargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetArraySize(val *float64) {
	if err := j.validateSetArraySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arraySize",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetInternalValue(val *TfTarget_BatchTargetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetJobAttempts(val *float64) {
	if err := j.validateSetJobAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobAttempts",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetJobDefinition(val *string) {
	if err := j.validateSetJobDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinition",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ResetArraySize() {
	_jsii_.InvokeVoid(
		t,
		"resetArraySize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ResetJobAttempts() {
	_jsii_.InvokeVoid(
		t,
		"resetJobAttempts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTarget_BatchTargetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

