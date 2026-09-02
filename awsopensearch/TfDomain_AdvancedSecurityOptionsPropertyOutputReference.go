package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_AdvancedSecurityOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AnonymousAuthEnabled() interface{}
	// Experimental.
	SetAnonymousAuthEnabled(val interface{})
	// Experimental.
	AnonymousAuthEnabledInput() interface{}
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalUserDatabaseEnabled() interface{}
	// Experimental.
	SetInternalUserDatabaseEnabled(val interface{})
	// Experimental.
	InternalUserDatabaseEnabledInput() interface{}
	// Experimental.
	InternalValue() *TfDomain_AdvancedSecurityOptionsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_AdvancedSecurityOptionsProperty)
	// Experimental.
	JwtOptions() TfDomain_JwtOptionsPropertyOutputReference
	// Experimental.
	JwtOptionsInput() *TfDomain_JwtOptionsProperty
	// Experimental.
	MasterUserOptions() TfDomain_MasterUserOptionsPropertyOutputReference
	// Experimental.
	MasterUserOptionsInput() *TfDomain_MasterUserOptionsProperty
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
	PutJwtOptions(value *TfDomain_JwtOptionsProperty)
	// Experimental.
	PutMasterUserOptions(value *TfDomain_MasterUserOptionsProperty)
	// Experimental.
	ResetAnonymousAuthEnabled()
	// Experimental.
	ResetInternalUserDatabaseEnabled()
	// Experimental.
	ResetJwtOptions()
	// Experimental.
	ResetMasterUserOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_AdvancedSecurityOptionsPropertyOutputReference
type jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) AnonymousAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) AnonymousAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) InternalUserDatabaseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) InternalUserDatabaseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) InternalValue() *TfDomain_AdvancedSecurityOptionsProperty {
	var returns *TfDomain_AdvancedSecurityOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) JwtOptions() TfDomain_JwtOptionsPropertyOutputReference {
	var returns TfDomain_JwtOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"jwtOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) JwtOptionsInput() *TfDomain_JwtOptionsProperty {
	var returns *TfDomain_JwtOptionsProperty
	_jsii_.Get(
		j,
		"jwtOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) MasterUserOptions() TfDomain_MasterUserOptionsPropertyOutputReference {
	var returns TfDomain_MasterUserOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"masterUserOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) MasterUserOptionsInput() *TfDomain_MasterUserOptionsProperty {
	var returns *TfDomain_MasterUserOptionsProperty
	_jsii_.Get(
		j,
		"masterUserOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_AdvancedSecurityOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_AdvancedSecurityOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_AdvancedSecurityOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.TfDomain.AdvancedSecurityOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_AdvancedSecurityOptionsPropertyOutputReference_Override(t TfDomain_AdvancedSecurityOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.TfDomain.AdvancedSecurityOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetAnonymousAuthEnabled(val interface{}) {
	if err := j.validateSetAnonymousAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetInternalUserDatabaseEnabled(val interface{}) {
	if err := j.validateSetInternalUserDatabaseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalUserDatabaseEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetInternalValue(val *TfDomain_AdvancedSecurityOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) PutJwtOptions(value *TfDomain_JwtOptionsProperty) {
	if err := t.validatePutJwtOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJwtOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) PutMasterUserOptions(value *TfDomain_MasterUserOptionsProperty) {
	if err := t.validatePutMasterUserOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMasterUserOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetAnonymousAuthEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetAnonymousAuthEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetInternalUserDatabaseEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetInternalUserDatabaseEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetJwtOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetJwtOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetMasterUserOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterUserOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_AdvancedSecurityOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

