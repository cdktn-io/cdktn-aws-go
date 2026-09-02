package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTask_ReportOverridesPropertyOutputReference interface {
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
	DeletedOverride() *string
	// Experimental.
	SetDeletedOverride(val *string)
	// Experimental.
	DeletedOverrideInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTask_ReportOverridesProperty
	// Experimental.
	SetInternalValue(val *TfTask_ReportOverridesProperty)
	// Experimental.
	SkippedOverride() *string
	// Experimental.
	SetSkippedOverride(val *string)
	// Experimental.
	SkippedOverrideInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransferredOverride() *string
	// Experimental.
	SetTransferredOverride(val *string)
	// Experimental.
	TransferredOverrideInput() *string
	// Experimental.
	VerifiedOverride() *string
	// Experimental.
	SetVerifiedOverride(val *string)
	// Experimental.
	VerifiedOverrideInput() *string
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
	ResetDeletedOverride()
	// Experimental.
	ResetSkippedOverride()
	// Experimental.
	ResetTransferredOverride()
	// Experimental.
	ResetVerifiedOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTask_ReportOverridesPropertyOutputReference
type jsiiProxy_TfTask_ReportOverridesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) DeletedOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletedOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) DeletedOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletedOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) InternalValue() *TfTask_ReportOverridesProperty {
	var returns *TfTask_ReportOverridesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) SkippedOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skippedOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) SkippedOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skippedOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) TransferredOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferredOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) TransferredOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferredOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) VerifiedOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) VerifiedOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedOverrideInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTask_ReportOverridesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTask_ReportOverridesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTask_ReportOverridesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTask_ReportOverridesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.TfTask.ReportOverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTask_ReportOverridesPropertyOutputReference_Override(t TfTask_ReportOverridesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.TfTask.ReportOverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetDeletedOverride(val *string) {
	if err := j.validateSetDeletedOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedOverride",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetInternalValue(val *TfTask_ReportOverridesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetSkippedOverride(val *string) {
	if err := j.validateSetSkippedOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skippedOverride",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetTransferredOverride(val *string) {
	if err := j.validateSetTransferredOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferredOverride",
		val,
	)
}

func (j *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference)SetVerifiedOverride(val *string) {
	if err := j.validateSetVerifiedOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedOverride",
		val,
	)
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ResetDeletedOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletedOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ResetSkippedOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetSkippedOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ResetTransferredOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetTransferredOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ResetVerifiedOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetVerifiedOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTask_ReportOverridesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

