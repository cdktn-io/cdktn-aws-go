package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_StudioWebPortalSettingsPropertyOutputReference interface {
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
	HiddenAppTypes() *[]*string
	// Experimental.
	SetHiddenAppTypes(val *[]*string)
	// Experimental.
	HiddenAppTypesInput() *[]*string
	// Experimental.
	HiddenInstanceTypes() *[]*string
	// Experimental.
	SetHiddenInstanceTypes(val *[]*string)
	// Experimental.
	HiddenInstanceTypesInput() *[]*string
	// Experimental.
	HiddenMlTools() *[]*string
	// Experimental.
	SetHiddenMlTools(val *[]*string)
	// Experimental.
	HiddenMlToolsInput() *[]*string
	// Experimental.
	InternalValue() *TfDomain_StudioWebPortalSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_StudioWebPortalSettingsProperty)
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
	ResetHiddenAppTypes()
	// Experimental.
	ResetHiddenInstanceTypes()
	// Experimental.
	ResetHiddenMlTools()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_StudioWebPortalSettingsPropertyOutputReference
type jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) HiddenAppTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenAppTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) HiddenAppTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenAppTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) HiddenInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) HiddenInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) HiddenMlTools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenMlTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) HiddenMlToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenMlToolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) InternalValue() *TfDomain_StudioWebPortalSettingsProperty {
	var returns *TfDomain_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_StudioWebPortalSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_StudioWebPortalSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_StudioWebPortalSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.StudioWebPortalSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_StudioWebPortalSettingsPropertyOutputReference_Override(t TfDomain_StudioWebPortalSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.StudioWebPortalSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetHiddenAppTypes(val *[]*string) {
	if err := j.validateSetHiddenAppTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenAppTypes",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetHiddenInstanceTypes(val *[]*string) {
	if err := j.validateSetHiddenInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetHiddenMlTools(val *[]*string) {
	if err := j.validateSetHiddenMlToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenMlTools",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_StudioWebPortalSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ResetHiddenAppTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetHiddenAppTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ResetHiddenInstanceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetHiddenInstanceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ResetHiddenMlTools() {
	_jsii_.InvokeVoid(
		t,
		"resetHiddenMlTools",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_StudioWebPortalSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

