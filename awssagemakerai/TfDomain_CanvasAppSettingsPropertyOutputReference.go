package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_CanvasAppSettingsPropertyOutputReference interface {
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
	DirectDeploySettings() TfDomain_DirectDeploySettingsPropertyOutputReference
	// Experimental.
	DirectDeploySettingsInput() *TfDomain_DirectDeploySettingsProperty
	// Experimental.
	EmrServerlessSettings() TfDomain_EmrServerlessSettingsPropertyOutputReference
	// Experimental.
	EmrServerlessSettingsInput() *TfDomain_EmrServerlessSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GenerativeAiSettings() TfDomain_GenerativeAiSettingsPropertyOutputReference
	// Experimental.
	GenerativeAiSettingsInput() *TfDomain_GenerativeAiSettingsProperty
	// Experimental.
	IdentityProviderOauthSettings() TfDomain_IdentityProviderOauthSettingsPropertyList
	// Experimental.
	IdentityProviderOauthSettingsInput() interface{}
	// Experimental.
	InternalValue() *TfDomain_CanvasAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_CanvasAppSettingsProperty)
	// Experimental.
	KendraSettings() TfDomain_KendraSettingsPropertyOutputReference
	// Experimental.
	KendraSettingsInput() *TfDomain_KendraSettingsProperty
	// Experimental.
	ModelRegisterSettings() TfDomain_ModelRegisterSettingsPropertyOutputReference
	// Experimental.
	ModelRegisterSettingsInput() *TfDomain_ModelRegisterSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeSeriesForecastingSettings() TfDomain_TimeSeriesForecastingSettingsPropertyOutputReference
	// Experimental.
	TimeSeriesForecastingSettingsInput() *TfDomain_TimeSeriesForecastingSettingsProperty
	// Experimental.
	WorkspaceSettings() TfDomain_WorkspaceSettingsPropertyOutputReference
	// Experimental.
	WorkspaceSettingsInput() *TfDomain_WorkspaceSettingsProperty
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
	PutDirectDeploySettings(value *TfDomain_DirectDeploySettingsProperty)
	// Experimental.
	PutEmrServerlessSettings(value *TfDomain_EmrServerlessSettingsProperty)
	// Experimental.
	PutGenerativeAiSettings(value *TfDomain_GenerativeAiSettingsProperty)
	// Experimental.
	PutIdentityProviderOauthSettings(value interface{})
	// Experimental.
	PutKendraSettings(value *TfDomain_KendraSettingsProperty)
	// Experimental.
	PutModelRegisterSettings(value *TfDomain_ModelRegisterSettingsProperty)
	// Experimental.
	PutTimeSeriesForecastingSettings(value *TfDomain_TimeSeriesForecastingSettingsProperty)
	// Experimental.
	PutWorkspaceSettings(value *TfDomain_WorkspaceSettingsProperty)
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

// The jsii proxy struct for TfDomain_CanvasAppSettingsPropertyOutputReference
type jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) DirectDeploySettings() TfDomain_DirectDeploySettingsPropertyOutputReference {
	var returns TfDomain_DirectDeploySettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"directDeploySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) DirectDeploySettingsInput() *TfDomain_DirectDeploySettingsProperty {
	var returns *TfDomain_DirectDeploySettingsProperty
	_jsii_.Get(
		j,
		"directDeploySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) EmrServerlessSettings() TfDomain_EmrServerlessSettingsPropertyOutputReference {
	var returns TfDomain_EmrServerlessSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrServerlessSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) EmrServerlessSettingsInput() *TfDomain_EmrServerlessSettingsProperty {
	var returns *TfDomain_EmrServerlessSettingsProperty
	_jsii_.Get(
		j,
		"emrServerlessSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GenerativeAiSettings() TfDomain_GenerativeAiSettingsPropertyOutputReference {
	var returns TfDomain_GenerativeAiSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"generativeAiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GenerativeAiSettingsInput() *TfDomain_GenerativeAiSettingsProperty {
	var returns *TfDomain_GenerativeAiSettingsProperty
	_jsii_.Get(
		j,
		"generativeAiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) IdentityProviderOauthSettings() TfDomain_IdentityProviderOauthSettingsPropertyList {
	var returns TfDomain_IdentityProviderOauthSettingsPropertyList
	_jsii_.Get(
		j,
		"identityProviderOauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) IdentityProviderOauthSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderOauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) InternalValue() *TfDomain_CanvasAppSettingsProperty {
	var returns *TfDomain_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) KendraSettings() TfDomain_KendraSettingsPropertyOutputReference {
	var returns TfDomain_KendraSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kendraSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) KendraSettingsInput() *TfDomain_KendraSettingsProperty {
	var returns *TfDomain_KendraSettingsProperty
	_jsii_.Get(
		j,
		"kendraSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ModelRegisterSettings() TfDomain_ModelRegisterSettingsPropertyOutputReference {
	var returns TfDomain_ModelRegisterSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"modelRegisterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ModelRegisterSettingsInput() *TfDomain_ModelRegisterSettingsProperty {
	var returns *TfDomain_ModelRegisterSettingsProperty
	_jsii_.Get(
		j,
		"modelRegisterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) TimeSeriesForecastingSettings() TfDomain_TimeSeriesForecastingSettingsPropertyOutputReference {
	var returns TfDomain_TimeSeriesForecastingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) TimeSeriesForecastingSettingsInput() *TfDomain_TimeSeriesForecastingSettingsProperty {
	var returns *TfDomain_TimeSeriesForecastingSettingsProperty
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) WorkspaceSettings() TfDomain_WorkspaceSettingsPropertyOutputReference {
	var returns TfDomain_WorkspaceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) WorkspaceSettingsInput() *TfDomain_WorkspaceSettingsProperty {
	var returns *TfDomain_WorkspaceSettingsProperty
	_jsii_.Get(
		j,
		"workspaceSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_CanvasAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_CanvasAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_CanvasAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.CanvasAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_CanvasAppSettingsPropertyOutputReference_Override(t TfDomain_CanvasAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.CanvasAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_CanvasAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutDirectDeploySettings(value *TfDomain_DirectDeploySettingsProperty) {
	if err := t.validatePutDirectDeploySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDirectDeploySettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutEmrServerlessSettings(value *TfDomain_EmrServerlessSettingsProperty) {
	if err := t.validatePutEmrServerlessSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmrServerlessSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutGenerativeAiSettings(value *TfDomain_GenerativeAiSettingsProperty) {
	if err := t.validatePutGenerativeAiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGenerativeAiSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutIdentityProviderOauthSettings(value interface{}) {
	if err := t.validatePutIdentityProviderOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIdentityProviderOauthSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutKendraSettings(value *TfDomain_KendraSettingsProperty) {
	if err := t.validatePutKendraSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKendraSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutModelRegisterSettings(value *TfDomain_ModelRegisterSettingsProperty) {
	if err := t.validatePutModelRegisterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelRegisterSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutTimeSeriesForecastingSettings(value *TfDomain_TimeSeriesForecastingSettingsProperty) {
	if err := t.validatePutTimeSeriesForecastingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeSeriesForecastingSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) PutWorkspaceSettings(value *TfDomain_WorkspaceSettingsProperty) {
	if err := t.validatePutWorkspaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkspaceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetDirectDeploySettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectDeploySettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetEmrServerlessSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmrServerlessSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetGenerativeAiSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetGenerativeAiSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetIdentityProviderOauthSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityProviderOauthSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetKendraSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKendraSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetModelRegisterSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetModelRegisterSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetTimeSeriesForecastingSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeSeriesForecastingSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ResetWorkspaceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_CanvasAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

