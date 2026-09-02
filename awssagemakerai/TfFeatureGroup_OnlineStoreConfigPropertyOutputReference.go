package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFeatureGroup_OnlineStoreConfigPropertyOutputReference interface {
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
	EnableOnlineStore() interface{}
	// Experimental.
	SetEnableOnlineStore(val interface{})
	// Experimental.
	EnableOnlineStoreInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFeatureGroup_OnlineStoreConfigProperty
	// Experimental.
	SetInternalValue(val *TfFeatureGroup_OnlineStoreConfigProperty)
	// Experimental.
	SecurityConfig() TfFeatureGroup_SecurityConfigPropertyOutputReference
	// Experimental.
	SecurityConfigInput() *TfFeatureGroup_SecurityConfigProperty
	// Experimental.
	StorageType() *string
	// Experimental.
	SetStorageType(val *string)
	// Experimental.
	StorageTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TtlDuration() TfFeatureGroup_TtlDurationPropertyOutputReference
	// Experimental.
	TtlDurationInput() *TfFeatureGroup_TtlDurationProperty
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
	PutSecurityConfig(value *TfFeatureGroup_SecurityConfigProperty)
	// Experimental.
	PutTtlDuration(value *TfFeatureGroup_TtlDurationProperty)
	// Experimental.
	ResetEnableOnlineStore()
	// Experimental.
	ResetSecurityConfig()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetTtlDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFeatureGroup_OnlineStoreConfigPropertyOutputReference
type jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) EnableOnlineStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnlineStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) EnableOnlineStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnlineStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) InternalValue() *TfFeatureGroup_OnlineStoreConfigProperty {
	var returns *TfFeatureGroup_OnlineStoreConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) SecurityConfig() TfFeatureGroup_SecurityConfigPropertyOutputReference {
	var returns TfFeatureGroup_SecurityConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) SecurityConfigInput() *TfFeatureGroup_SecurityConfigProperty {
	var returns *TfFeatureGroup_SecurityConfigProperty
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) TtlDuration() TfFeatureGroup_TtlDurationPropertyOutputReference {
	var returns TfFeatureGroup_TtlDurationPropertyOutputReference
	_jsii_.Get(
		j,
		"ttlDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) TtlDurationInput() *TfFeatureGroup_TtlDurationProperty {
	var returns *TfFeatureGroup_TtlDurationProperty
	_jsii_.Get(
		j,
		"ttlDurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFeatureGroup_OnlineStoreConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFeatureGroup_OnlineStoreConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFeatureGroup_OnlineStoreConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFeatureGroup.OnlineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFeatureGroup_OnlineStoreConfigPropertyOutputReference_Override(t TfFeatureGroup_OnlineStoreConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFeatureGroup.OnlineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetEnableOnlineStore(val interface{}) {
	if err := j.validateSetEnableOnlineStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOnlineStore",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetInternalValue(val *TfFeatureGroup_OnlineStoreConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) PutSecurityConfig(value *TfFeatureGroup_SecurityConfigProperty) {
	if err := t.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) PutTtlDuration(value *TfFeatureGroup_TtlDurationProperty) {
	if err := t.validatePutTtlDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTtlDuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetEnableOnlineStore() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableOnlineStore",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetTtlDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetTtlDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFeatureGroup_OnlineStoreConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

