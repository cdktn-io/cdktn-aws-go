package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkflow_StepsPropertyOutputReference interface {
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
	CopyStepDetails() TfWorkflow_StepsCopyStepDetailsPropertyOutputReference
	// Experimental.
	CopyStepDetailsInput() *TfWorkflow_StepsCopyStepDetailsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomStepDetails() TfWorkflow_StepsCustomStepDetailsPropertyOutputReference
	// Experimental.
	CustomStepDetailsInput() *TfWorkflow_StepsCustomStepDetailsProperty
	// Experimental.
	DecryptStepDetails() TfWorkflow_StepsDecryptStepDetailsPropertyOutputReference
	// Experimental.
	DecryptStepDetailsInput() *TfWorkflow_StepsDecryptStepDetailsProperty
	// Experimental.
	DeleteStepDetails() TfWorkflow_StepsDeleteStepDetailsPropertyOutputReference
	// Experimental.
	DeleteStepDetailsInput() *TfWorkflow_StepsDeleteStepDetailsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TagStepDetails() TfWorkflow_StepsTagStepDetailsPropertyOutputReference
	// Experimental.
	TagStepDetailsInput() *TfWorkflow_StepsTagStepDetailsProperty
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
	PutCopyStepDetails(value *TfWorkflow_StepsCopyStepDetailsProperty)
	// Experimental.
	PutCustomStepDetails(value *TfWorkflow_StepsCustomStepDetailsProperty)
	// Experimental.
	PutDecryptStepDetails(value *TfWorkflow_StepsDecryptStepDetailsProperty)
	// Experimental.
	PutDeleteStepDetails(value *TfWorkflow_StepsDeleteStepDetailsProperty)
	// Experimental.
	PutTagStepDetails(value *TfWorkflow_StepsTagStepDetailsProperty)
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

// The jsii proxy struct for TfWorkflow_StepsPropertyOutputReference
type jsiiProxy_TfWorkflow_StepsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) CopyStepDetails() TfWorkflow_StepsCopyStepDetailsPropertyOutputReference {
	var returns TfWorkflow_StepsCopyStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"copyStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) CopyStepDetailsInput() *TfWorkflow_StepsCopyStepDetailsProperty {
	var returns *TfWorkflow_StepsCopyStepDetailsProperty
	_jsii_.Get(
		j,
		"copyStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) CustomStepDetails() TfWorkflow_StepsCustomStepDetailsPropertyOutputReference {
	var returns TfWorkflow_StepsCustomStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"customStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) CustomStepDetailsInput() *TfWorkflow_StepsCustomStepDetailsProperty {
	var returns *TfWorkflow_StepsCustomStepDetailsProperty
	_jsii_.Get(
		j,
		"customStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) DecryptStepDetails() TfWorkflow_StepsDecryptStepDetailsPropertyOutputReference {
	var returns TfWorkflow_StepsDecryptStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"decryptStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) DecryptStepDetailsInput() *TfWorkflow_StepsDecryptStepDetailsProperty {
	var returns *TfWorkflow_StepsDecryptStepDetailsProperty
	_jsii_.Get(
		j,
		"decryptStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) DeleteStepDetails() TfWorkflow_StepsDeleteStepDetailsPropertyOutputReference {
	var returns TfWorkflow_StepsDeleteStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) DeleteStepDetailsInput() *TfWorkflow_StepsDeleteStepDetailsProperty {
	var returns *TfWorkflow_StepsDeleteStepDetailsProperty
	_jsii_.Get(
		j,
		"deleteStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) TagStepDetails() TfWorkflow_StepsTagStepDetailsPropertyOutputReference {
	var returns TfWorkflow_StepsTagStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"tagStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) TagStepDetailsInput() *TfWorkflow_StepsTagStepDetailsProperty {
	var returns *TfWorkflow_StepsTagStepDetailsProperty
	_jsii_.Get(
		j,
		"tagStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkflow_StepsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWorkflow_StepsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkflow_StepsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkflow_StepsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.StepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkflow_StepsPropertyOutputReference_Override(t TfWorkflow_StepsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.StepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) PutCopyStepDetails(value *TfWorkflow_StepsCopyStepDetailsProperty) {
	if err := t.validatePutCopyStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCopyStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) PutCustomStepDetails(value *TfWorkflow_StepsCustomStepDetailsProperty) {
	if err := t.validatePutCustomStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) PutDecryptStepDetails(value *TfWorkflow_StepsDecryptStepDetailsProperty) {
	if err := t.validatePutDecryptStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDecryptStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) PutDeleteStepDetails(value *TfWorkflow_StepsDeleteStepDetailsProperty) {
	if err := t.validatePutDeleteStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeleteStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) PutTagStepDetails(value *TfWorkflow_StepsTagStepDetailsProperty) {
	if err := t.validatePutTagStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ResetCopyStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ResetCustomStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ResetDecryptStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDecryptStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ResetDeleteStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ResetTagStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetTagStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkflow_StepsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

