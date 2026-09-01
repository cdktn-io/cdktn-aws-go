package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDmsEndpoint_MysqlSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsDmsEndpoint_MysqlSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDmsEndpoint_MysqlSettingsProperty)
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

// The jsii proxy struct for AwsDmsEndpoint_MysqlSettingsPropertyOutputReference
type jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) CleanSourceMetadataOnMismatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) CleanSourceMetadataOnMismatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) EventsPollInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) EventsPollIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) InternalValue() *AwsDmsEndpoint_MysqlSettingsProperty {
	var returns *AwsDmsEndpoint_MysqlSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ParallelLoadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ParallelLoadThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ServerTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ServerTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) TargetDbType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) TargetDbTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDmsEndpoint_MysqlSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDmsEndpoint_MysqlSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDmsEndpoint_MysqlSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsDmsEndpoint.MysqlSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDmsEndpoint_MysqlSettingsPropertyOutputReference_Override(a AwsDmsEndpoint_MysqlSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsDmsEndpoint.MysqlSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetCleanSourceMetadataOnMismatch(val interface{}) {
	if err := j.validateSetCleanSourceMetadataOnMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanSourceMetadataOnMismatch",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetEventsPollInterval(val *float64) {
	if err := j.validateSetEventsPollIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsPollInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetInternalValue(val *AwsDmsEndpoint_MysqlSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetParallelLoadThreads(val *float64) {
	if err := j.validateSetParallelLoadThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelLoadThreads",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetServerTimezone(val *string) {
	if err := j.validateSetServerTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTimezone",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetTargetDbType(val *string) {
	if err := j.validateSetTargetDbTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbType",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		a,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetCleanSourceMetadataOnMismatch() {
	_jsii_.InvokeVoid(
		a,
		"resetCleanSourceMetadataOnMismatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetEventsPollInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetEventsPollInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetParallelLoadThreads() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelLoadThreads",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetServerTimezone() {
	_jsii_.InvokeVoid(
		a,
		"resetServerTimezone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ResetTargetDbType() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetDbType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDmsEndpoint_MysqlSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

