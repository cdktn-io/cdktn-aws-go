package awselastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipeline_NotificationsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Completed() *string
	// Experimental.
	SetCompleted(val *string)
	// Experimental.
	CompletedInput() *string
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
	Error() *string
	// Experimental.
	SetError(val *string)
	// Experimental.
	ErrorInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPipeline_NotificationsProperty
	// Experimental.
	SetInternalValue(val *TfPipeline_NotificationsProperty)
	// Experimental.
	Progressing() *string
	// Experimental.
	SetProgressing(val *string)
	// Experimental.
	ProgressingInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Warning() *string
	// Experimental.
	SetWarning(val *string)
	// Experimental.
	WarningInput() *string
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
	ResetCompleted()
	// Experimental.
	ResetError()
	// Experimental.
	ResetProgressing()
	// Experimental.
	ResetWarning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipeline_NotificationsPropertyOutputReference
type jsiiProxy_TfPipeline_NotificationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) Completed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) CompletedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) Error() *string {
	var returns *string
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ErrorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) InternalValue() *TfPipeline_NotificationsProperty {
	var returns *TfPipeline_NotificationsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) Progressing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"progressing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ProgressingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"progressingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) WarningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipeline_NotificationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipeline_NotificationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipeline_NotificationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipeline_NotificationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.TfPipeline.NotificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipeline_NotificationsPropertyOutputReference_Override(t TfPipeline_NotificationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.TfPipeline.NotificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetCompleted(val *string) {
	if err := j.validateSetCompletedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completed",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetError(val *string) {
	if err := j.validateSetErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"error",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetInternalValue(val *TfPipeline_NotificationsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetProgressing(val *string) {
	if err := j.validateSetProgressingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"progressing",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference)SetWarning(val *string) {
	if err := j.validateSetWarningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warning",
		val,
	)
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ResetCompleted() {
	_jsii_.InvokeVoid(
		t,
		"resetCompleted",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ResetError() {
	_jsii_.InvokeVoid(
		t,
		"resetError",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ResetProgressing() {
	_jsii_.InvokeVoid(
		t,
		"resetProgressing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		t,
		"resetWarning",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipeline_NotificationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

