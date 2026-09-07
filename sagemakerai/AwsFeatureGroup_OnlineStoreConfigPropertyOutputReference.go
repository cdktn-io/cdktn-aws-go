package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference interface {
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
	InternalValue() *AwsFeatureGroup_OnlineStoreConfigProperty
	// Experimental.
	SetInternalValue(val *AwsFeatureGroup_OnlineStoreConfigProperty)
	// Experimental.
	SecurityConfig() AwsFeatureGroup_SecurityConfigPropertyOutputReference
	// Experimental.
	SecurityConfigInput() *AwsFeatureGroup_SecurityConfigProperty
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
	TtlDuration() AwsFeatureGroup_TtlDurationPropertyOutputReference
	// Experimental.
	TtlDurationInput() *AwsFeatureGroup_TtlDurationProperty
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
	PutSecurityConfig(value *AwsFeatureGroup_SecurityConfigProperty)
	// Experimental.
	PutTtlDuration(value *AwsFeatureGroup_TtlDurationProperty)
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

// The jsii proxy struct for AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference
type jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) EnableOnlineStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnlineStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) EnableOnlineStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnlineStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) InternalValue() *AwsFeatureGroup_OnlineStoreConfigProperty {
	var returns *AwsFeatureGroup_OnlineStoreConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) SecurityConfig() AwsFeatureGroup_SecurityConfigPropertyOutputReference {
	var returns AwsFeatureGroup_SecurityConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) SecurityConfigInput() *AwsFeatureGroup_SecurityConfigProperty {
	var returns *AwsFeatureGroup_SecurityConfigProperty
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) TtlDuration() AwsFeatureGroup_TtlDurationPropertyOutputReference {
	var returns AwsFeatureGroup_TtlDurationPropertyOutputReference
	_jsii_.Get(
		j,
		"ttlDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) TtlDurationInput() *AwsFeatureGroup_TtlDurationProperty {
	var returns *AwsFeatureGroup_TtlDurationProperty
	_jsii_.Get(
		j,
		"ttlDurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFeatureGroup_OnlineStoreConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFeatureGroup_OnlineStoreConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsFeatureGroup.OnlineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFeatureGroup_OnlineStoreConfigPropertyOutputReference_Override(a AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsFeatureGroup.OnlineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetEnableOnlineStore(val interface{}) {
	if err := j.validateSetEnableOnlineStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOnlineStore",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetInternalValue(val *AwsFeatureGroup_OnlineStoreConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) PutSecurityConfig(value *AwsFeatureGroup_SecurityConfigProperty) {
	if err := a.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) PutTtlDuration(value *AwsFeatureGroup_TtlDurationProperty) {
	if err := a.validatePutTtlDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTtlDuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetEnableOnlineStore() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableOnlineStore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetStorageType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ResetTtlDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetTtlDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFeatureGroup_OnlineStoreConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

