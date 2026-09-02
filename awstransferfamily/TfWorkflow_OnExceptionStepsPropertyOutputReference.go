package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkflow_OnExceptionStepsPropertyOutputReference interface {
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
	// Experimental.
	CopyStepDetails() TfWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference
	// Experimental.
	CopyStepDetailsInput() *TfWorkflow_OnExceptionStepsCopyStepDetailsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomStepDetails() TfWorkflow_OnExceptionStepsCustomStepDetailsPropertyOutputReference
	// Experimental.
	CustomStepDetailsInput() *TfWorkflow_OnExceptionStepsCustomStepDetailsProperty
	// Experimental.
	DecryptStepDetails() TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference
	// Experimental.
	DecryptStepDetailsInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty
	// Experimental.
	DeleteStepDetails() TfWorkflow_OnExceptionStepsDeleteStepDetailsPropertyOutputReference
	// Experimental.
	DeleteStepDetailsInput() *TfWorkflow_OnExceptionStepsDeleteStepDetailsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TagStepDetails() TfWorkflow_OnExceptionStepsTagStepDetailsPropertyOutputReference
	// Experimental.
	TagStepDetailsInput() *TfWorkflow_OnExceptionStepsTagStepDetailsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutCopyStepDetails(value *TfWorkflow_OnExceptionStepsCopyStepDetailsProperty)
	// Experimental.
	PutCustomStepDetails(value *TfWorkflow_OnExceptionStepsCustomStepDetailsProperty)
	// Experimental.
	PutDecryptStepDetails(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty)
	// Experimental.
	PutDeleteStepDetails(value *TfWorkflow_OnExceptionStepsDeleteStepDetailsProperty)
	// Experimental.
	PutTagStepDetails(value *TfWorkflow_OnExceptionStepsTagStepDetailsProperty)
	// Experimental.
	ResetCopyStepDetails()
	// Experimental.
	ResetCustomStepDetails()
	// Experimental.
	ResetDecryptStepDetails()
	// Experimental.
	ResetDeleteStepDetails()
	// Experimental.
	ResetTagStepDetails()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWorkflow_OnExceptionStepsPropertyOutputReference
type jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) CopyStepDetails() TfWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"copyStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) CopyStepDetailsInput() *TfWorkflow_OnExceptionStepsCopyStepDetailsProperty {
	var returns *TfWorkflow_OnExceptionStepsCopyStepDetailsProperty
	_jsii_.Get(
		j,
		"copyStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) CustomStepDetails() TfWorkflow_OnExceptionStepsCustomStepDetailsPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsCustomStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"customStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) CustomStepDetailsInput() *TfWorkflow_OnExceptionStepsCustomStepDetailsProperty {
	var returns *TfWorkflow_OnExceptionStepsCustomStepDetailsProperty
	_jsii_.Get(
		j,
		"customStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) DecryptStepDetails() TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"decryptStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) DecryptStepDetailsInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty {
	var returns *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty
	_jsii_.Get(
		j,
		"decryptStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) DeleteStepDetails() TfWorkflow_OnExceptionStepsDeleteStepDetailsPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsDeleteStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) DeleteStepDetailsInput() *TfWorkflow_OnExceptionStepsDeleteStepDetailsProperty {
	var returns *TfWorkflow_OnExceptionStepsDeleteStepDetailsProperty
	_jsii_.Get(
		j,
		"deleteStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) TagStepDetails() TfWorkflow_OnExceptionStepsTagStepDetailsPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsTagStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"tagStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) TagStepDetailsInput() *TfWorkflow_OnExceptionStepsTagStepDetailsProperty {
	var returns *TfWorkflow_OnExceptionStepsTagStepDetailsProperty
	_jsii_.Get(
		j,
		"tagStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkflow_OnExceptionStepsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWorkflow_OnExceptionStepsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkflow_OnExceptionStepsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkflow_OnExceptionStepsPropertyOutputReference_Override(t TfWorkflow_OnExceptionStepsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) PutCopyStepDetails(value *TfWorkflow_OnExceptionStepsCopyStepDetailsProperty) {
	if err := t.validatePutCopyStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCopyStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) PutCustomStepDetails(value *TfWorkflow_OnExceptionStepsCustomStepDetailsProperty) {
	if err := t.validatePutCustomStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) PutDecryptStepDetails(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty) {
	if err := t.validatePutDecryptStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDecryptStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) PutDeleteStepDetails(value *TfWorkflow_OnExceptionStepsDeleteStepDetailsProperty) {
	if err := t.validatePutDeleteStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeleteStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) PutTagStepDetails(value *TfWorkflow_OnExceptionStepsTagStepDetailsProperty) {
	if err := t.validatePutTagStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ResetCopyStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ResetCustomStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ResetDecryptStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDecryptStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ResetDeleteStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ResetTagStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetTagStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

