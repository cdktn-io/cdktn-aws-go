package awscloudwatchrum

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchrum/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchrum/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAppMonitor_AppMonitorConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowCookies() interface{}
	// Experimental.
	SetAllowCookies(val interface{})
	// Experimental.
	AllowCookiesInput() interface{}
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
	EnableXray() interface{}
	// Experimental.
	SetEnableXray(val interface{})
	// Experimental.
	EnableXrayInput() interface{}
	// Experimental.
	ExcludedPages() *[]*string
	// Experimental.
	SetExcludedPages(val *[]*string)
	// Experimental.
	ExcludedPagesInput() *[]*string
	// Experimental.
	FavoritePages() *[]*string
	// Experimental.
	SetFavoritePages(val *[]*string)
	// Experimental.
	FavoritePagesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	GuestRoleArn() *string
	// Experimental.
	SetGuestRoleArn(val *string)
	// Experimental.
	GuestRoleArnInput() *string
	// Experimental.
	IdentityPoolId() *string
	// Experimental.
	SetIdentityPoolId(val *string)
	// Experimental.
	IdentityPoolIdInput() *string
	// Experimental.
	IncludedPages() *[]*string
	// Experimental.
	SetIncludedPages(val *[]*string)
	// Experimental.
	IncludedPagesInput() *[]*string
	// Experimental.
	InternalValue() *TfAppMonitor_AppMonitorConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfAppMonitor_AppMonitorConfigurationProperty)
	// Experimental.
	SessionSampleRate() *float64
	// Experimental.
	SetSessionSampleRate(val *float64)
	// Experimental.
	SessionSampleRateInput() *float64
	// Experimental.
	Telemetries() *[]*string
	// Experimental.
	SetTelemetries(val *[]*string)
	// Experimental.
	TelemetriesInput() *[]*string
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
	ResetAllowCookies()
	// Experimental.
	ResetEnableXray()
	// Experimental.
	ResetExcludedPages()
	// Experimental.
	ResetFavoritePages()
	// Experimental.
	ResetGuestRoleArn()
	// Experimental.
	ResetIdentityPoolId()
	// Experimental.
	ResetIncludedPages()
	// Experimental.
	ResetSessionSampleRate()
	// Experimental.
	ResetTelemetries()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAppMonitor_AppMonitorConfigurationPropertyOutputReference
type jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) AllowCookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) AllowCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) EnableXray() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) EnableXrayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXrayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ExcludedPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ExcludedPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) FavoritePages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"favoritePages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) FavoritePagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"favoritePagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GuestRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GuestRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) IncludedPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) IncludedPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) InternalValue() *TfAppMonitor_AppMonitorConfigurationProperty {
	var returns *TfAppMonitor_AppMonitorConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) SessionSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) SessionSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) Telemetries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) TelemetriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAppMonitor_AppMonitorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAppMonitor_AppMonitorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAppMonitor_AppMonitorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-rum.TfAppMonitor.AppMonitorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAppMonitor_AppMonitorConfigurationPropertyOutputReference_Override(t TfAppMonitor_AppMonitorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-rum.TfAppMonitor.AppMonitorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetAllowCookies(val interface{}) {
	if err := j.validateSetAllowCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCookies",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetEnableXray(val interface{}) {
	if err := j.validateSetEnableXrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXray",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetExcludedPages(val *[]*string) {
	if err := j.validateSetExcludedPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPages",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetFavoritePages(val *[]*string) {
	if err := j.validateSetFavoritePagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"favoritePages",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetGuestRoleArn(val *string) {
	if err := j.validateSetGuestRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetIdentityPoolId(val *string) {
	if err := j.validateSetIdentityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetIncludedPages(val *[]*string) {
	if err := j.validateSetIncludedPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedPages",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetInternalValue(val *TfAppMonitor_AppMonitorConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetSessionSampleRate(val *float64) {
	if err := j.validateSetSessionSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionSampleRate",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetTelemetries(val *[]*string) {
	if err := j.validateSetTelemetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetries",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetAllowCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetEnableXray() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableXray",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetExcludedPages() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludedPages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetFavoritePages() {
	_jsii_.InvokeVoid(
		t,
		"resetFavoritePages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetGuestRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetGuestRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetIdentityPoolId() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityPoolId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetIncludedPages() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludedPages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetSessionSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetTelemetries() {
	_jsii_.InvokeVoid(
		t,
		"resetTelemetries",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAppMonitor_AppMonitorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

