package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference interface {
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
	DefaultEbsVolumeSizeInGb() *float64
	// Experimental.
	SetDefaultEbsVolumeSizeInGb(val *float64)
	// Experimental.
	DefaultEbsVolumeSizeInGbInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty)
	// Experimental.
	MaximumEbsVolumeSizeInGb() *float64
	// Experimental.
	SetMaximumEbsVolumeSizeInGb(val *float64)
	// Experimental.
	MaximumEbsVolumeSizeInGbInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference
type jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) DefaultEbsVolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultEbsVolumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) DefaultEbsVolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultEbsVolumeSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) InternalValue() *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty {
	var returns *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) MaximumEbsVolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEbsVolumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) MaximumEbsVolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEbsVolumeSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference_Override(t TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetDefaultEbsVolumeSizeInGb(val *float64) {
	if err := j.validateSetDefaultEbsVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEbsVolumeSizeInGb",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetMaximumEbsVolumeSizeInGb(val *float64) {
	if err := j.validateSetMaximumEbsVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumEbsVolumeSizeInGb",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

