package elasticsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elasticsearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elasticsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomainSamlOptions_SamlOptionsPropertyOutputReference interface {
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Idp() AwsDomainSamlOptions_IdpPropertyOutputReference
	// Experimental.
	IdpInput() *AwsDomainSamlOptions_IdpProperty
	// Experimental.
	InternalValue() *AwsDomainSamlOptions_SamlOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsDomainSamlOptions_SamlOptionsProperty)
	// Experimental.
	MasterBackendRole() *string
	// Experimental.
	SetMasterBackendRole(val *string)
	// Experimental.
	MasterBackendRoleInput() *string
	// Experimental.
	MasterUserName() *string
	// Experimental.
	SetMasterUserName(val *string)
	// Experimental.
	MasterUserNameInput() *string
	// Experimental.
	RolesKey() *string
	// Experimental.
	SetRolesKey(val *string)
	// Experimental.
	RolesKeyInput() *string
	// Experimental.
	SessionTimeoutMinutes() *float64
	// Experimental.
	SetSessionTimeoutMinutes(val *float64)
	// Experimental.
	SessionTimeoutMinutesInput() *float64
	// Experimental.
	SubjectKey() *string
	// Experimental.
	SetSubjectKey(val *string)
	// Experimental.
	SubjectKeyInput() *string
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
	PutIdp(value *AwsDomainSamlOptions_IdpProperty)
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetIdp()
	// Experimental.
	ResetMasterBackendRole()
	// Experimental.
	ResetMasterUserName()
	// Experimental.
	ResetRolesKey()
	// Experimental.
	ResetSessionTimeoutMinutes()
	// Experimental.
	ResetSubjectKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDomainSamlOptions_SamlOptionsPropertyOutputReference
type jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) Idp() AwsDomainSamlOptions_IdpPropertyOutputReference {
	var returns AwsDomainSamlOptions_IdpPropertyOutputReference
	_jsii_.Get(
		j,
		"idp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) IdpInput() *AwsDomainSamlOptions_IdpProperty {
	var returns *AwsDomainSamlOptions_IdpProperty
	_jsii_.Get(
		j,
		"idpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) InternalValue() *AwsDomainSamlOptions_SamlOptionsProperty {
	var returns *AwsDomainSamlOptions_SamlOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) MasterBackendRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterBackendRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) MasterBackendRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterBackendRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) MasterUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) MasterUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) SessionTimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) SessionTimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomainSamlOptions_SamlOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomainSamlOptions_SamlOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomainSamlOptions_SamlOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elasticsearch.AwsDomainSamlOptions.SamlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomainSamlOptions_SamlOptionsPropertyOutputReference_Override(a AwsDomainSamlOptions_SamlOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticsearch.AwsDomainSamlOptions.SamlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetInternalValue(val *AwsDomainSamlOptions_SamlOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetMasterBackendRole(val *string) {
	if err := j.validateSetMasterBackendRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterBackendRole",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetMasterUserName(val *string) {
	if err := j.validateSetMasterUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserName",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetRolesKey(val *string) {
	if err := j.validateSetRolesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetSessionTimeoutMinutes(val *float64) {
	if err := j.validateSetSessionTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetSubjectKey(val *string) {
	if err := j.validateSetSubjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) PutIdp(value *AwsDomainSamlOptions_IdpProperty) {
	if err := a.validatePutIdpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetIdp() {
	_jsii_.InvokeVoid(
		a,
		"resetIdp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetMasterBackendRole() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterBackendRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetMasterUserName() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterUserName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetSessionTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomainSamlOptions_SamlOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

