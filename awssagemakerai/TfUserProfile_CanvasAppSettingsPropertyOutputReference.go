package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_CanvasAppSettingsPropertyOutputReference interface {
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
	DirectDeploySettings() TfUserProfile_DirectDeploySettingsPropertyOutputReference
	// Experimental.
	DirectDeploySettingsInput() *TfUserProfile_DirectDeploySettingsProperty
	// Experimental.
	EmrServerlessSettings() TfUserProfile_EmrServerlessSettingsPropertyOutputReference
	// Experimental.
	EmrServerlessSettingsInput() *TfUserProfile_EmrServerlessSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GenerativeAiSettings() TfUserProfile_GenerativeAiSettingsPropertyOutputReference
	// Experimental.
	GenerativeAiSettingsInput() *TfUserProfile_GenerativeAiSettingsProperty
	// Experimental.
	IdentityProviderOauthSettings() TfUserProfile_IdentityProviderOauthSettingsPropertyList
	// Experimental.
	IdentityProviderOauthSettingsInput() interface{}
	// Experimental.
	InternalValue() *TfUserProfile_CanvasAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_CanvasAppSettingsProperty)
	// Experimental.
	KendraSettings() TfUserProfile_KendraSettingsPropertyOutputReference
	// Experimental.
	KendraSettingsInput() *TfUserProfile_KendraSettingsProperty
	// Experimental.
	ModelRegisterSettings() TfUserProfile_ModelRegisterSettingsPropertyOutputReference
	// Experimental.
	ModelRegisterSettingsInput() *TfUserProfile_ModelRegisterSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeSeriesForecastingSettings() TfUserProfile_TimeSeriesForecastingSettingsPropertyOutputReference
	// Experimental.
	TimeSeriesForecastingSettingsInput() *TfUserProfile_TimeSeriesForecastingSettingsProperty
	// Experimental.
	WorkspaceSettings() TfUserProfile_WorkspaceSettingsPropertyOutputReference
	// Experimental.
	WorkspaceSettingsInput() *TfUserProfile_WorkspaceSettingsProperty
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
	PutDirectDeploySettings(value *TfUserProfile_DirectDeploySettingsProperty)
	// Experimental.
	PutEmrServerlessSettings(value *TfUserProfile_EmrServerlessSettingsProperty)
	// Experimental.
	PutGenerativeAiSettings(value *TfUserProfile_GenerativeAiSettingsProperty)
	// Experimental.
	PutIdentityProviderOauthSettings(value interface{})
	// Experimental.
	PutKendraSettings(value *TfUserProfile_KendraSettingsProperty)
	// Experimental.
	PutModelRegisterSettings(value *TfUserProfile_ModelRegisterSettingsProperty)
	// Experimental.
	PutTimeSeriesForecastingSettings(value *TfUserProfile_TimeSeriesForecastingSettingsProperty)
	// Experimental.
	PutWorkspaceSettings(value *TfUserProfile_WorkspaceSettingsProperty)
	// Experimental.
	ResetDirectDeploySettings()
	// Experimental.
	ResetEmrServerlessSettings()
	// Experimental.
	ResetGenerativeAiSettings()
	// Experimental.
	ResetIdentityProviderOauthSettings()
	// Experimental.
	ResetKendraSettings()
	// Experimental.
	ResetModelRegisterSettings()
	// Experimental.
	ResetTimeSeriesForecastingSettings()
	// Experimental.
	ResetWorkspaceSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserProfile_CanvasAppSettingsPropertyOutputReference
type jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) DirectDeploySettings() TfUserProfile_DirectDeploySettingsPropertyOutputReference {
	var returns TfUserProfile_DirectDeploySettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"directDeploySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) DirectDeploySettingsInput() *TfUserProfile_DirectDeploySettingsProperty {
	var returns *TfUserProfile_DirectDeploySettingsProperty
	_jsii_.Get(
		j,
		"directDeploySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) EmrServerlessSettings() TfUserProfile_EmrServerlessSettingsPropertyOutputReference {
	var returns TfUserProfile_EmrServerlessSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrServerlessSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) EmrServerlessSettingsInput() *TfUserProfile_EmrServerlessSettingsProperty {
	var returns *TfUserProfile_EmrServerlessSettingsProperty
	_jsii_.Get(
		j,
		"emrServerlessSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GenerativeAiSettings() TfUserProfile_GenerativeAiSettingsPropertyOutputReference {
	var returns TfUserProfile_GenerativeAiSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"generativeAiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GenerativeAiSettingsInput() *TfUserProfile_GenerativeAiSettingsProperty {
	var returns *TfUserProfile_GenerativeAiSettingsProperty
	_jsii_.Get(
		j,
		"generativeAiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) IdentityProviderOauthSettings() TfUserProfile_IdentityProviderOauthSettingsPropertyList {
	var returns TfUserProfile_IdentityProviderOauthSettingsPropertyList
	_jsii_.Get(
		j,
		"identityProviderOauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) IdentityProviderOauthSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderOauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) InternalValue() *TfUserProfile_CanvasAppSettingsProperty {
	var returns *TfUserProfile_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) KendraSettings() TfUserProfile_KendraSettingsPropertyOutputReference {
	var returns TfUserProfile_KendraSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kendraSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) KendraSettingsInput() *TfUserProfile_KendraSettingsProperty {
	var returns *TfUserProfile_KendraSettingsProperty
	_jsii_.Get(
		j,
		"kendraSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ModelRegisterSettings() TfUserProfile_ModelRegisterSettingsPropertyOutputReference {
	var returns TfUserProfile_ModelRegisterSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"modelRegisterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ModelRegisterSettingsInput() *TfUserProfile_ModelRegisterSettingsProperty {
	var returns *TfUserProfile_ModelRegisterSettingsProperty
	_jsii_.Get(
		j,
		"modelRegisterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) TimeSeriesForecastingSettings() TfUserProfile_TimeSeriesForecastingSettingsPropertyOutputReference {
	var returns TfUserProfile_TimeSeriesForecastingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) TimeSeriesForecastingSettingsInput() *TfUserProfile_TimeSeriesForecastingSettingsProperty {
	var returns *TfUserProfile_TimeSeriesForecastingSettingsProperty
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) WorkspaceSettings() TfUserProfile_WorkspaceSettingsPropertyOutputReference {
	var returns TfUserProfile_WorkspaceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) WorkspaceSettingsInput() *TfUserProfile_WorkspaceSettingsProperty {
	var returns *TfUserProfile_WorkspaceSettingsProperty
	_jsii_.Get(
		j,
		"workspaceSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_CanvasAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_CanvasAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_CanvasAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.CanvasAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_CanvasAppSettingsPropertyOutputReference_Override(t TfUserProfile_CanvasAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.CanvasAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference)SetInternalValue(val *TfUserProfile_CanvasAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutDirectDeploySettings(value *TfUserProfile_DirectDeploySettingsProperty) {
	if err := t.validatePutDirectDeploySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDirectDeploySettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutEmrServerlessSettings(value *TfUserProfile_EmrServerlessSettingsProperty) {
	if err := t.validatePutEmrServerlessSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmrServerlessSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutGenerativeAiSettings(value *TfUserProfile_GenerativeAiSettingsProperty) {
	if err := t.validatePutGenerativeAiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGenerativeAiSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutIdentityProviderOauthSettings(value interface{}) {
	if err := t.validatePutIdentityProviderOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIdentityProviderOauthSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutKendraSettings(value *TfUserProfile_KendraSettingsProperty) {
	if err := t.validatePutKendraSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKendraSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutModelRegisterSettings(value *TfUserProfile_ModelRegisterSettingsProperty) {
	if err := t.validatePutModelRegisterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelRegisterSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutTimeSeriesForecastingSettings(value *TfUserProfile_TimeSeriesForecastingSettingsProperty) {
	if err := t.validatePutTimeSeriesForecastingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeSeriesForecastingSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) PutWorkspaceSettings(value *TfUserProfile_WorkspaceSettingsProperty) {
	if err := t.validatePutWorkspaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkspaceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetDirectDeploySettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectDeploySettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetEmrServerlessSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmrServerlessSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetGenerativeAiSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetGenerativeAiSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetIdentityProviderOauthSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityProviderOauthSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetKendraSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKendraSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetModelRegisterSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetModelRegisterSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetTimeSeriesForecastingSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeSeriesForecastingSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ResetWorkspaceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_CanvasAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

