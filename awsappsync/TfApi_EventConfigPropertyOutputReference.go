package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApi_EventConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthProvider() TfApi_AuthProviderPropertyList
	// Experimental.
	AuthProviderInput() interface{}
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
	ConnectionAuthMode() TfApi_ConnectionAuthModePropertyList
	// Experimental.
	ConnectionAuthModeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DefaultPublishAuthMode() TfApi_DefaultPublishAuthModePropertyList
	// Experimental.
	DefaultPublishAuthModeInput() interface{}
	// Experimental.
	DefaultSubscribeAuthMode() TfApi_DefaultSubscribeAuthModePropertyList
	// Experimental.
	DefaultSubscribeAuthModeInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogConfig() TfApi_LogConfigPropertyList
	// Experimental.
	LogConfigInput() interface{}
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
	PutAuthProvider(value interface{})
	// Experimental.
	PutConnectionAuthMode(value interface{})
	// Experimental.
	PutDefaultPublishAuthMode(value interface{})
	// Experimental.
	PutDefaultSubscribeAuthMode(value interface{})
	// Experimental.
	PutLogConfig(value interface{})
	// Experimental.
	ResetAuthProvider()
	// Experimental.
	ResetConnectionAuthMode()
	// Experimental.
	ResetDefaultPublishAuthMode()
	// Experimental.
	ResetDefaultSubscribeAuthMode()
	// Experimental.
	ResetLogConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApi_EventConfigPropertyOutputReference
type jsiiProxy_TfApi_EventConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) AuthProvider() TfApi_AuthProviderPropertyList {
	var returns TfApi_AuthProviderPropertyList
	_jsii_.Get(
		j,
		"authProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) AuthProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ConnectionAuthMode() TfApi_ConnectionAuthModePropertyList {
	var returns TfApi_ConnectionAuthModePropertyList
	_jsii_.Get(
		j,
		"connectionAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ConnectionAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) DefaultPublishAuthMode() TfApi_DefaultPublishAuthModePropertyList {
	var returns TfApi_DefaultPublishAuthModePropertyList
	_jsii_.Get(
		j,
		"defaultPublishAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) DefaultPublishAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPublishAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) DefaultSubscribeAuthMode() TfApi_DefaultSubscribeAuthModePropertyList {
	var returns TfApi_DefaultSubscribeAuthModePropertyList
	_jsii_.Get(
		j,
		"defaultSubscribeAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) DefaultSubscribeAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSubscribeAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) LogConfig() TfApi_LogConfigPropertyList {
	var returns TfApi_LogConfigPropertyList
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) LogConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApi_EventConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfApi_EventConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApi_EventConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApi_EventConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.TfApi.EventConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApi_EventConfigPropertyOutputReference_Override(t TfApi_EventConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.TfApi.EventConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApi_EventConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) PutAuthProvider(value interface{}) {
	if err := t.validatePutAuthProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthProvider",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) PutConnectionAuthMode(value interface{}) {
	if err := t.validatePutConnectionAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionAuthMode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) PutDefaultPublishAuthMode(value interface{}) {
	if err := t.validatePutDefaultPublishAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultPublishAuthMode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) PutDefaultSubscribeAuthMode(value interface{}) {
	if err := t.validatePutDefaultSubscribeAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultSubscribeAuthMode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) PutLogConfig(value interface{}) {
	if err := t.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ResetAuthProvider() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthProvider",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ResetConnectionAuthMode() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionAuthMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ResetDefaultPublishAuthMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultPublishAuthMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ResetDefaultSubscribeAuthMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultSubscribeAuthMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ResetLogConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApi_EventConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

