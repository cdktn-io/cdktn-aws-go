package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_CanvasAppSettingsPropertyOutputReference interface {
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
	DirectDeploySettings() AwsDomain_DirectDeploySettingsPropertyOutputReference
	// Experimental.
	DirectDeploySettingsInput() *AwsDomain_DirectDeploySettingsProperty
	// Experimental.
	EmrServerlessSettings() AwsDomain_EmrServerlessSettingsPropertyOutputReference
	// Experimental.
	EmrServerlessSettingsInput() *AwsDomain_EmrServerlessSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GenerativeAiSettings() AwsDomain_GenerativeAiSettingsPropertyOutputReference
	// Experimental.
	GenerativeAiSettingsInput() *AwsDomain_GenerativeAiSettingsProperty
	// Experimental.
	IdentityProviderOauthSettings() AwsDomain_IdentityProviderOauthSettingsPropertyList
	// Experimental.
	IdentityProviderOauthSettingsInput() interface{}
	// Experimental.
	InternalValue() *AwsDomain_CanvasAppSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_CanvasAppSettingsProperty)
	// Experimental.
	KendraSettings() AwsDomain_KendraSettingsPropertyOutputReference
	// Experimental.
	KendraSettingsInput() *AwsDomain_KendraSettingsProperty
	// Experimental.
	ModelRegisterSettings() AwsDomain_ModelRegisterSettingsPropertyOutputReference
	// Experimental.
	ModelRegisterSettingsInput() *AwsDomain_ModelRegisterSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeSeriesForecastingSettings() AwsDomain_TimeSeriesForecastingSettingsPropertyOutputReference
	// Experimental.
	TimeSeriesForecastingSettingsInput() *AwsDomain_TimeSeriesForecastingSettingsProperty
	// Experimental.
	WorkspaceSettings() AwsDomain_WorkspaceSettingsPropertyOutputReference
	// Experimental.
	WorkspaceSettingsInput() *AwsDomain_WorkspaceSettingsProperty
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
	PutDirectDeploySettings(value *AwsDomain_DirectDeploySettingsProperty)
	// Experimental.
	PutEmrServerlessSettings(value *AwsDomain_EmrServerlessSettingsProperty)
	// Experimental.
	PutGenerativeAiSettings(value *AwsDomain_GenerativeAiSettingsProperty)
	// Experimental.
	PutIdentityProviderOauthSettings(value interface{})
	// Experimental.
	PutKendraSettings(value *AwsDomain_KendraSettingsProperty)
	// Experimental.
	PutModelRegisterSettings(value *AwsDomain_ModelRegisterSettingsProperty)
	// Experimental.
	PutTimeSeriesForecastingSettings(value *AwsDomain_TimeSeriesForecastingSettingsProperty)
	// Experimental.
	PutWorkspaceSettings(value *AwsDomain_WorkspaceSettingsProperty)
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

// The jsii proxy struct for AwsDomain_CanvasAppSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) DirectDeploySettings() AwsDomain_DirectDeploySettingsPropertyOutputReference {
	var returns AwsDomain_DirectDeploySettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"directDeploySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) DirectDeploySettingsInput() *AwsDomain_DirectDeploySettingsProperty {
	var returns *AwsDomain_DirectDeploySettingsProperty
	_jsii_.Get(
		j,
		"directDeploySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) EmrServerlessSettings() AwsDomain_EmrServerlessSettingsPropertyOutputReference {
	var returns AwsDomain_EmrServerlessSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrServerlessSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) EmrServerlessSettingsInput() *AwsDomain_EmrServerlessSettingsProperty {
	var returns *AwsDomain_EmrServerlessSettingsProperty
	_jsii_.Get(
		j,
		"emrServerlessSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GenerativeAiSettings() AwsDomain_GenerativeAiSettingsPropertyOutputReference {
	var returns AwsDomain_GenerativeAiSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"generativeAiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GenerativeAiSettingsInput() *AwsDomain_GenerativeAiSettingsProperty {
	var returns *AwsDomain_GenerativeAiSettingsProperty
	_jsii_.Get(
		j,
		"generativeAiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) IdentityProviderOauthSettings() AwsDomain_IdentityProviderOauthSettingsPropertyList {
	var returns AwsDomain_IdentityProviderOauthSettingsPropertyList
	_jsii_.Get(
		j,
		"identityProviderOauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) IdentityProviderOauthSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderOauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) InternalValue() *AwsDomain_CanvasAppSettingsProperty {
	var returns *AwsDomain_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) KendraSettings() AwsDomain_KendraSettingsPropertyOutputReference {
	var returns AwsDomain_KendraSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kendraSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) KendraSettingsInput() *AwsDomain_KendraSettingsProperty {
	var returns *AwsDomain_KendraSettingsProperty
	_jsii_.Get(
		j,
		"kendraSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ModelRegisterSettings() AwsDomain_ModelRegisterSettingsPropertyOutputReference {
	var returns AwsDomain_ModelRegisterSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"modelRegisterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ModelRegisterSettingsInput() *AwsDomain_ModelRegisterSettingsProperty {
	var returns *AwsDomain_ModelRegisterSettingsProperty
	_jsii_.Get(
		j,
		"modelRegisterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) TimeSeriesForecastingSettings() AwsDomain_TimeSeriesForecastingSettingsPropertyOutputReference {
	var returns AwsDomain_TimeSeriesForecastingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) TimeSeriesForecastingSettingsInput() *AwsDomain_TimeSeriesForecastingSettingsProperty {
	var returns *AwsDomain_TimeSeriesForecastingSettingsProperty
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) WorkspaceSettings() AwsDomain_WorkspaceSettingsPropertyOutputReference {
	var returns AwsDomain_WorkspaceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) WorkspaceSettingsInput() *AwsDomain_WorkspaceSettingsProperty {
	var returns *AwsDomain_WorkspaceSettingsProperty
	_jsii_.Get(
		j,
		"workspaceSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_CanvasAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_CanvasAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_CanvasAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.CanvasAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_CanvasAppSettingsPropertyOutputReference_Override(a AwsDomain_CanvasAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.CanvasAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_CanvasAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutDirectDeploySettings(value *AwsDomain_DirectDeploySettingsProperty) {
	if err := a.validatePutDirectDeploySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDirectDeploySettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutEmrServerlessSettings(value *AwsDomain_EmrServerlessSettingsProperty) {
	if err := a.validatePutEmrServerlessSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmrServerlessSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutGenerativeAiSettings(value *AwsDomain_GenerativeAiSettingsProperty) {
	if err := a.validatePutGenerativeAiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGenerativeAiSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutIdentityProviderOauthSettings(value interface{}) {
	if err := a.validatePutIdentityProviderOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentityProviderOauthSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutKendraSettings(value *AwsDomain_KendraSettingsProperty) {
	if err := a.validatePutKendraSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKendraSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutModelRegisterSettings(value *AwsDomain_ModelRegisterSettingsProperty) {
	if err := a.validatePutModelRegisterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putModelRegisterSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutTimeSeriesForecastingSettings(value *AwsDomain_TimeSeriesForecastingSettingsProperty) {
	if err := a.validatePutTimeSeriesForecastingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeSeriesForecastingSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) PutWorkspaceSettings(value *AwsDomain_WorkspaceSettingsProperty) {
	if err := a.validatePutWorkspaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetDirectDeploySettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectDeploySettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetEmrServerlessSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmrServerlessSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetGenerativeAiSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerativeAiSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetIdentityProviderOauthSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityProviderOauthSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetKendraSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKendraSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetModelRegisterSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetModelRegisterSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetTimeSeriesForecastingSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeSeriesForecastingSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ResetWorkspaceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_CanvasAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

