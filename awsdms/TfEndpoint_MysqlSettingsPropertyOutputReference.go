package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_MysqlSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AfterConnectScript() *string
	// Experimental.
	SetAfterConnectScript(val *string)
	// Experimental.
	AfterConnectScriptInput() *string
	// Experimental.
	AuthenticationMethod() *string
	// Experimental.
	SetAuthenticationMethod(val *string)
	// Experimental.
	AuthenticationMethodInput() *string
	// Experimental.
	CleanSourceMetadataOnMismatch() interface{}
	// Experimental.
	SetCleanSourceMetadataOnMismatch(val interface{})
	// Experimental.
	CleanSourceMetadataOnMismatchInput() interface{}
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
	EventsPollInterval() *float64
	// Experimental.
	SetEventsPollInterval(val *float64)
	// Experimental.
	EventsPollIntervalInput() *float64
	// Experimental.
	ExecuteTimeout() *float64
	// Experimental.
	SetExecuteTimeout(val *float64)
	// Experimental.
	ExecuteTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEndpoint_MysqlSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_MysqlSettingsProperty)
	// Experimental.
	MaxFileSize() *float64
	// Experimental.
	SetMaxFileSize(val *float64)
	// Experimental.
	MaxFileSizeInput() *float64
	// Experimental.
	ParallelLoadThreads() *float64
	// Experimental.
	SetParallelLoadThreads(val *float64)
	// Experimental.
	ParallelLoadThreadsInput() *float64
	// Experimental.
	ServerTimezone() *string
	// Experimental.
	SetServerTimezone(val *string)
	// Experimental.
	ServerTimezoneInput() *string
	// Experimental.
	ServiceAccessRoleArn() *string
	// Experimental.
	SetServiceAccessRoleArn(val *string)
	// Experimental.
	ServiceAccessRoleArnInput() *string
	// Experimental.
	TargetDbType() *string
	// Experimental.
	SetTargetDbType(val *string)
	// Experimental.
	TargetDbTypeInput() *string
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
	ResetAfterConnectScript()
	// Experimental.
	ResetAuthenticationMethod()
	// Experimental.
	ResetCleanSourceMetadataOnMismatch()
	// Experimental.
	ResetEventsPollInterval()
	// Experimental.
	ResetExecuteTimeout()
	// Experimental.
	ResetMaxFileSize()
	// Experimental.
	ResetParallelLoadThreads()
	// Experimental.
	ResetServerTimezone()
	// Experimental.
	ResetServiceAccessRoleArn()
	// Experimental.
	ResetTargetDbType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_MysqlSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) CleanSourceMetadataOnMismatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) CleanSourceMetadataOnMismatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) EventsPollInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) EventsPollIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) InternalValue() *TfEndpoint_MysqlSettingsProperty {
	var returns *TfEndpoint_MysqlSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ParallelLoadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ParallelLoadThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ServerTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ServerTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) TargetDbType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) TargetDbTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_MysqlSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_MysqlSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_MysqlSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.MysqlSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_MysqlSettingsPropertyOutputReference_Override(t TfEndpoint_MysqlSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.MysqlSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetCleanSourceMetadataOnMismatch(val interface{}) {
	if err := j.validateSetCleanSourceMetadataOnMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanSourceMetadataOnMismatch",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetEventsPollInterval(val *float64) {
	if err := j.validateSetEventsPollIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsPollInterval",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_MysqlSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetParallelLoadThreads(val *float64) {
	if err := j.validateSetParallelLoadThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelLoadThreads",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetServerTimezone(val *string) {
	if err := j.validateSetServerTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTimezone",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetTargetDbType(val *string) {
	if err := j.validateSetTargetDbTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbType",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		t,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetCleanSourceMetadataOnMismatch() {
	_jsii_.InvokeVoid(
		t,
		"resetCleanSourceMetadataOnMismatch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetEventsPollInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetEventsPollInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetParallelLoadThreads() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelLoadThreads",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetServerTimezone() {
	_jsii_.InvokeVoid(
		t,
		"resetServerTimezone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ResetTargetDbType() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetDbType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_MysqlSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

