package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCustomActionType_SettingsPropertyOutputReference interface {
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
	EntityUrlTemplate() *string
	// Experimental.
	SetEntityUrlTemplate(val *string)
	// Experimental.
	EntityUrlTemplateInput() *string
	// Experimental.
	ExecutionUrlTemplate() *string
	// Experimental.
	SetExecutionUrlTemplate(val *string)
	// Experimental.
	ExecutionUrlTemplateInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCustomActionType_SettingsProperty
	// Experimental.
	SetInternalValue(val *TfCustomActionType_SettingsProperty)
	// Experimental.
	RevisionUrlTemplate() *string
	// Experimental.
	SetRevisionUrlTemplate(val *string)
	// Experimental.
	RevisionUrlTemplateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThirdPartyConfigurationUrl() *string
	// Experimental.
	SetThirdPartyConfigurationUrl(val *string)
	// Experimental.
	ThirdPartyConfigurationUrlInput() *string
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
	ResetEntityUrlTemplate()
	// Experimental.
	ResetExecutionUrlTemplate()
	// Experimental.
	ResetRevisionUrlTemplate()
	// Experimental.
	ResetThirdPartyConfigurationUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCustomActionType_SettingsPropertyOutputReference
type jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) EntityUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) EntityUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ExecutionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ExecutionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) InternalValue() *TfCustomActionType_SettingsProperty {
	var returns *TfCustomActionType_SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) RevisionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) RevisionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ThirdPartyConfigurationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thirdPartyConfigurationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ThirdPartyConfigurationUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thirdPartyConfigurationUrlInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCustomActionType_SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCustomActionType_SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCustomActionType_SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCustomActionType.SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCustomActionType_SettingsPropertyOutputReference_Override(t TfCustomActionType_SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCustomActionType.SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetEntityUrlTemplate(val *string) {
	if err := j.validateSetEntityUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetExecutionUrlTemplate(val *string) {
	if err := j.validateSetExecutionUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetInternalValue(val *TfCustomActionType_SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetRevisionUrlTemplate(val *string) {
	if err := j.validateSetRevisionUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference)SetThirdPartyConfigurationUrl(val *string) {
	if err := j.validateSetThirdPartyConfigurationUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thirdPartyConfigurationUrl",
		val,
	)
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ResetEntityUrlTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetEntityUrlTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ResetExecutionUrlTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionUrlTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ResetRevisionUrlTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetRevisionUrlTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ResetThirdPartyConfigurationUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetThirdPartyConfigurationUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCustomActionType_SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

