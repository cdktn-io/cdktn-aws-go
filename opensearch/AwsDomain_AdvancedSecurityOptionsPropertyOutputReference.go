package opensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/opensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/opensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_AdvancedSecurityOptionsPropertyOutputReference interface {
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
	InternalValue() *AwsDomain_AdvancedSecurityOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_AdvancedSecurityOptionsProperty)
	// Experimental.
	JwtOptions() AwsDomain_JwtOptionsPropertyOutputReference
	// Experimental.
	JwtOptionsInput() *AwsDomain_JwtOptionsProperty
	// Experimental.
	MasterUserOptions() AwsDomain_MasterUserOptionsPropertyOutputReference
	// Experimental.
	MasterUserOptionsInput() *AwsDomain_MasterUserOptionsProperty
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
	PutJwtOptions(value *AwsDomain_JwtOptionsProperty)
	// Experimental.
	PutMasterUserOptions(value *AwsDomain_MasterUserOptionsProperty)
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

// The jsii proxy struct for AwsDomain_AdvancedSecurityOptionsPropertyOutputReference
type jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) AnonymousAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) AnonymousAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) InternalUserDatabaseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) InternalUserDatabaseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) InternalValue() *AwsDomain_AdvancedSecurityOptionsProperty {
	var returns *AwsDomain_AdvancedSecurityOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) JwtOptions() AwsDomain_JwtOptionsPropertyOutputReference {
	var returns AwsDomain_JwtOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"jwtOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) JwtOptionsInput() *AwsDomain_JwtOptionsProperty {
	var returns *AwsDomain_JwtOptionsProperty
	_jsii_.Get(
		j,
		"jwtOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) MasterUserOptions() AwsDomain_MasterUserOptionsPropertyOutputReference {
	var returns AwsDomain_MasterUserOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"masterUserOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) MasterUserOptionsInput() *AwsDomain_MasterUserOptionsProperty {
	var returns *AwsDomain_MasterUserOptionsProperty
	_jsii_.Get(
		j,
		"masterUserOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_AdvancedSecurityOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_AdvancedSecurityOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_AdvancedSecurityOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsDomain.AdvancedSecurityOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_AdvancedSecurityOptionsPropertyOutputReference_Override(a AwsDomain_AdvancedSecurityOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsDomain.AdvancedSecurityOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetAnonymousAuthEnabled(val interface{}) {
	if err := j.validateSetAnonymousAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetInternalUserDatabaseEnabled(val interface{}) {
	if err := j.validateSetInternalUserDatabaseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalUserDatabaseEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetInternalValue(val *AwsDomain_AdvancedSecurityOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) PutJwtOptions(value *AwsDomain_JwtOptionsProperty) {
	if err := a.validatePutJwtOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJwtOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) PutMasterUserOptions(value *AwsDomain_MasterUserOptionsProperty) {
	if err := a.validatePutMasterUserOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMasterUserOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetAnonymousAuthEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAnonymousAuthEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetInternalUserDatabaseEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInternalUserDatabaseEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetJwtOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetJwtOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ResetMasterUserOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterUserOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_AdvancedSecurityOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

