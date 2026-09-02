package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointConfiguration_NotificationConfigPropertyOutputReference interface {
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
	ErrorTopic() *string
	// Experimental.
	SetErrorTopic(val *string)
	// Experimental.
	ErrorTopicInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeInferenceResponseIn() *[]*string
	// Experimental.
	SetIncludeInferenceResponseIn(val *[]*string)
	// Experimental.
	IncludeInferenceResponseInInput() *[]*string
	// Experimental.
	InternalValue() *TfEndpointConfiguration_NotificationConfigProperty
	// Experimental.
	SetInternalValue(val *TfEndpointConfiguration_NotificationConfigProperty)
	// Experimental.
	SuccessTopic() *string
	// Experimental.
	SetSuccessTopic(val *string)
	// Experimental.
	SuccessTopicInput() *string
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
	ResetErrorTopic()
	// Experimental.
	ResetIncludeInferenceResponseIn()
	// Experimental.
	ResetSuccessTopic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpointConfiguration_NotificationConfigPropertyOutputReference
type jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ErrorTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ErrorTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) IncludeInferenceResponseIn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInferenceResponseIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) IncludeInferenceResponseInInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInferenceResponseInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) InternalValue() *TfEndpointConfiguration_NotificationConfigProperty {
	var returns *TfEndpointConfiguration_NotificationConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) SuccessTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) SuccessTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpointConfiguration_NotificationConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpointConfiguration_NotificationConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpointConfiguration_NotificationConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.NotificationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpointConfiguration_NotificationConfigPropertyOutputReference_Override(t TfEndpointConfiguration_NotificationConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.NotificationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetErrorTopic(val *string) {
	if err := j.validateSetErrorTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorTopic",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetIncludeInferenceResponseIn(val *[]*string) {
	if err := j.validateSetIncludeInferenceResponseInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInferenceResponseIn",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetInternalValue(val *TfEndpointConfiguration_NotificationConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetSuccessTopic(val *string) {
	if err := j.validateSetSuccessTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successTopic",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ResetErrorTopic() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorTopic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ResetIncludeInferenceResponseIn() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeInferenceResponseIn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ResetSuccessTopic() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessTopic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_NotificationConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

