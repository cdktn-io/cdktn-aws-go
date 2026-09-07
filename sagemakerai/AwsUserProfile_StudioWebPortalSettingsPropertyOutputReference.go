package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsUserProfile_StudioWebPortalSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsUserProfile_StudioWebPortalSettingsProperty)
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

// The jsii proxy struct for AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference
type jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) HiddenAppTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenAppTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) HiddenAppTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenAppTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) HiddenInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) HiddenInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) HiddenMlTools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenMlTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) HiddenMlToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenMlToolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) InternalValue() *AwsUserProfile_StudioWebPortalSettingsProperty {
	var returns *AwsUserProfile_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUserProfile_StudioWebPortalSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUserProfile_StudioWebPortalSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.StudioWebPortalSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUserProfile_StudioWebPortalSettingsPropertyOutputReference_Override(a AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.StudioWebPortalSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetHiddenAppTypes(val *[]*string) {
	if err := j.validateSetHiddenAppTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenAppTypes",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetHiddenInstanceTypes(val *[]*string) {
	if err := j.validateSetHiddenInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetHiddenMlTools(val *[]*string) {
	if err := j.validateSetHiddenMlToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenMlTools",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetInternalValue(val *AwsUserProfile_StudioWebPortalSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ResetHiddenAppTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetHiddenAppTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ResetHiddenInstanceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetHiddenInstanceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ResetHiddenMlTools() {
	_jsii_.InvokeVoid(
		a,
		"resetHiddenMlTools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

