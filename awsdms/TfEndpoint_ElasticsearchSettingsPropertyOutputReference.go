package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_ElasticsearchSettingsPropertyOutputReference interface {
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
	EndpointUri() *string
	// Experimental.
	SetEndpointUri(val *string)
	// Experimental.
	EndpointUriInput() *string
	// Experimental.
	ErrorRetryDuration() *float64
	// Experimental.
	SetErrorRetryDuration(val *float64)
	// Experimental.
	ErrorRetryDurationInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	FullLoadErrorPercentage() *float64
	// Experimental.
	SetFullLoadErrorPercentage(val *float64)
	// Experimental.
	FullLoadErrorPercentageInput() *float64
	// Experimental.
	InternalValue() *TfEndpoint_ElasticsearchSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_ElasticsearchSettingsProperty)
	// Experimental.
	ServiceAccessRoleArn() *string
	// Experimental.
	SetServiceAccessRoleArn(val *string)
	// Experimental.
	ServiceAccessRoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseNewMappingType() interface{}
	// Experimental.
	SetUseNewMappingType(val interface{})
	// Experimental.
	UseNewMappingTypeInput() interface{}
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
	ResetErrorRetryDuration()
	// Experimental.
	ResetFullLoadErrorPercentage()
	// Experimental.
	ResetUseNewMappingType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_ElasticsearchSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) EndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ErrorRetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ErrorRetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) FullLoadErrorPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullLoadErrorPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) FullLoadErrorPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullLoadErrorPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) InternalValue() *TfEndpoint_ElasticsearchSettingsProperty {
	var returns *TfEndpoint_ElasticsearchSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) UseNewMappingType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNewMappingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) UseNewMappingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNewMappingTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_ElasticsearchSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_ElasticsearchSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_ElasticsearchSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.ElasticsearchSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_ElasticsearchSettingsPropertyOutputReference_Override(t TfEndpoint_ElasticsearchSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.ElasticsearchSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetEndpointUri(val *string) {
	if err := j.validateSetEndpointUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUri",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetErrorRetryDuration(val *float64) {
	if err := j.validateSetErrorRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorRetryDuration",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetFullLoadErrorPercentage(val *float64) {
	if err := j.validateSetFullLoadErrorPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullLoadErrorPercentage",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_ElasticsearchSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference)SetUseNewMappingType(val interface{}) {
	if err := j.validateSetUseNewMappingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNewMappingType",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ResetErrorRetryDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorRetryDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ResetFullLoadErrorPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetFullLoadErrorPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ResetUseNewMappingType() {
	_jsii_.InvokeVoid(
		t,
		"resetUseNewMappingType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_ElasticsearchSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

