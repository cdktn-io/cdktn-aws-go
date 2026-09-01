package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference interface {
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
	InternalValue() *AwsSagemakerEndpointConfiguration_NotificationConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerEndpointConfiguration_NotificationConfigProperty)
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

// The jsii proxy struct for AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ErrorTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ErrorTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) IncludeInferenceResponseIn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInferenceResponseIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) IncludeInferenceResponseInInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInferenceResponseInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) InternalValue() *AwsSagemakerEndpointConfiguration_NotificationConfigProperty {
	var returns *AwsSagemakerEndpointConfiguration_NotificationConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) SuccessTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) SuccessTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpointConfiguration.NotificationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference_Override(a AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpointConfiguration.NotificationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetErrorTopic(val *string) {
	if err := j.validateSetErrorTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorTopic",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetIncludeInferenceResponseIn(val *[]*string) {
	if err := j.validateSetIncludeInferenceResponseInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInferenceResponseIn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerEndpointConfiguration_NotificationConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetSuccessTopic(val *string) {
	if err := j.validateSetSuccessTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successTopic",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ResetErrorTopic() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorTopic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ResetIncludeInferenceResponseIn() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeInferenceResponseIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ResetSuccessTopic() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessTopic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_NotificationConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

