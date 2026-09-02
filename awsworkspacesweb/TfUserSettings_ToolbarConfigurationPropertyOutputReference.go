package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspacesweb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsworkspacesweb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserSettings_ToolbarConfigurationPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	HiddenToolbarItems() *[]*string
	// Experimental.
	SetHiddenToolbarItems(val *[]*string)
	// Experimental.
	HiddenToolbarItemsInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxDisplayResolution() *string
	// Experimental.
	SetMaxDisplayResolution(val *string)
	// Experimental.
	MaxDisplayResolutionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ToolbarType() *string
	// Experimental.
	SetToolbarType(val *string)
	// Experimental.
	ToolbarTypeInput() *string
	// Experimental.
	VisualMode() *string
	// Experimental.
	SetVisualMode(val *string)
	// Experimental.
	VisualModeInput() *string
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
	ResetHiddenToolbarItems()
	// Experimental.
	ResetMaxDisplayResolution()
	// Experimental.
	ResetToolbarType()
	// Experimental.
	ResetVisualMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserSettings_ToolbarConfigurationPropertyOutputReference
type jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) HiddenToolbarItems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenToolbarItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) HiddenToolbarItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenToolbarItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) MaxDisplayResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDisplayResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) MaxDisplayResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDisplayResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ToolbarType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolbarType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ToolbarTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolbarTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) VisualMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) VisualModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserSettings_ToolbarConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfUserSettings_ToolbarConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserSettings_ToolbarConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.TfUserSettings.ToolbarConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserSettings_ToolbarConfigurationPropertyOutputReference_Override(t TfUserSettings_ToolbarConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.TfUserSettings.ToolbarConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetHiddenToolbarItems(val *[]*string) {
	if err := j.validateSetHiddenToolbarItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenToolbarItems",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetMaxDisplayResolution(val *string) {
	if err := j.validateSetMaxDisplayResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDisplayResolution",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetToolbarType(val *string) {
	if err := j.validateSetToolbarTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolbarType",
		val,
	)
}

func (j *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference)SetVisualMode(val *string) {
	if err := j.validateSetVisualModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualMode",
		val,
	)
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ResetHiddenToolbarItems() {
	_jsii_.InvokeVoid(
		t,
		"resetHiddenToolbarItems",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ResetMaxDisplayResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxDisplayResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ResetToolbarType() {
	_jsii_.InvokeVoid(
		t,
		"resetToolbarType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ResetVisualMode() {
	_jsii_.InvokeVoid(
		t,
		"resetVisualMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserSettings_ToolbarConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

