package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference interface {
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
	DestinationFileLocation() TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference
	// Experimental.
	DestinationFileLocationInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty
	// Experimental.
	SetInternalValue(val *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	OverwriteExisting() *string
	// Experimental.
	SetOverwriteExisting(val *string)
	// Experimental.
	OverwriteExistingInput() *string
	// Experimental.
	SourceFileLocation() *string
	// Experimental.
	SetSourceFileLocation(val *string)
	// Experimental.
	SourceFileLocationInput() *string
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
	PutDestinationFileLocation(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty)
	// Experimental.
	ResetDestinationFileLocation()
	// Experimental.
	ResetName()
	// Experimental.
	ResetOverwriteExisting()
	// Experimental.
	ResetSourceFileLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference
type jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) DestinationFileLocation() TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"destinationFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) DestinationFileLocationInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty {
	var returns *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty
	_jsii_.Get(
		j,
		"destinationFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) InternalValue() *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty {
	var returns *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) OverwriteExisting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteExisting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) OverwriteExistingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteExistingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) SourceFileLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) SourceFileLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsDecryptStepDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference_Override(t TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsDecryptStepDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetInternalValue(val *TfWorkflow_OnExceptionStepsDecryptStepDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetOverwriteExisting(val *string) {
	if err := j.validateSetOverwriteExistingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteExisting",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetSourceFileLocation(val *string) {
	if err := j.validateSetSourceFileLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFileLocation",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) PutDestinationFileLocation(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty) {
	if err := t.validatePutDestinationFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationFileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ResetDestinationFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ResetOverwriteExisting() {
	_jsii_.InvokeVoid(
		t,
		"resetOverwriteExisting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ResetSourceFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

