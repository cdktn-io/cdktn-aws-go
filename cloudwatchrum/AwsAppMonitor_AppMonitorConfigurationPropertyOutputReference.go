package cloudwatchrum

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudwatchrum/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudwatchrum/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsAppMonitor_AppMonitorConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsAppMonitor_AppMonitorConfigurationProperty)
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

// The jsii proxy struct for AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference
type jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) AllowCookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) AllowCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) EnableXray() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) EnableXrayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXrayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ExcludedPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ExcludedPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) FavoritePages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"favoritePages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) FavoritePagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"favoritePagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GuestRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GuestRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) IncludedPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) IncludedPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) InternalValue() *AwsAppMonitor_AppMonitorConfigurationProperty {
	var returns *AwsAppMonitor_AppMonitorConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) SessionSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) SessionSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) Telemetries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) TelemetriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppMonitor_AppMonitorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppMonitor_AppMonitorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-rum.AwsAppMonitor.AppMonitorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppMonitor_AppMonitorConfigurationPropertyOutputReference_Override(a AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-rum.AwsAppMonitor.AppMonitorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetAllowCookies(val interface{}) {
	if err := j.validateSetAllowCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCookies",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetEnableXray(val interface{}) {
	if err := j.validateSetEnableXrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXray",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetExcludedPages(val *[]*string) {
	if err := j.validateSetExcludedPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPages",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetFavoritePages(val *[]*string) {
	if err := j.validateSetFavoritePagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"favoritePages",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetGuestRoleArn(val *string) {
	if err := j.validateSetGuestRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetIdentityPoolId(val *string) {
	if err := j.validateSetIdentityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetIncludedPages(val *[]*string) {
	if err := j.validateSetIncludedPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedPages",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetInternalValue(val *AwsAppMonitor_AppMonitorConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetSessionSampleRate(val *float64) {
	if err := j.validateSetSessionSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetTelemetries(val *[]*string) {
	if err := j.validateSetTelemetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetries",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetAllowCookies() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowCookies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetEnableXray() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableXray",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetExcludedPages() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedPages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetFavoritePages() {
	_jsii_.InvokeVoid(
		a,
		"resetFavoritePages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetGuestRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetGuestRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetIdentityPoolId() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityPoolId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetIncludedPages() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludedPages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetSessionSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ResetTelemetries() {
	_jsii_.InvokeVoid(
		a,
		"resetTelemetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppMonitor_AppMonitorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

