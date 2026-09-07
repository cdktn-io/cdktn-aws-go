package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_ElasticsearchSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsEndpoint_ElasticsearchSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_ElasticsearchSettingsProperty)
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

// The jsii proxy struct for AwsEndpoint_ElasticsearchSettingsPropertyOutputReference
type jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) EndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ErrorRetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ErrorRetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) FullLoadErrorPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullLoadErrorPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) FullLoadErrorPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullLoadErrorPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) InternalValue() *AwsEndpoint_ElasticsearchSettingsProperty {
	var returns *AwsEndpoint_ElasticsearchSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) UseNewMappingType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNewMappingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) UseNewMappingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNewMappingTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_ElasticsearchSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_ElasticsearchSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_ElasticsearchSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.ElasticsearchSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_ElasticsearchSettingsPropertyOutputReference_Override(a AwsEndpoint_ElasticsearchSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.ElasticsearchSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetEndpointUri(val *string) {
	if err := j.validateSetEndpointUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUri",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetErrorRetryDuration(val *float64) {
	if err := j.validateSetErrorRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorRetryDuration",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetFullLoadErrorPercentage(val *float64) {
	if err := j.validateSetFullLoadErrorPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullLoadErrorPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetInternalValue(val *AwsEndpoint_ElasticsearchSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference)SetUseNewMappingType(val interface{}) {
	if err := j.validateSetUseNewMappingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNewMappingType",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ResetErrorRetryDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorRetryDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ResetFullLoadErrorPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetFullLoadErrorPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ResetUseNewMappingType() {
	_jsii_.InvokeVoid(
		a,
		"resetUseNewMappingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_ElasticsearchSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

