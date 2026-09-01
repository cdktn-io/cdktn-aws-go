package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference interface {
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
	CopyStepDetails() AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference
	// Experimental.
	CopyStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomStepDetails() AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsPropertyOutputReference
	// Experimental.
	CustomStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsProperty
	// Experimental.
	DecryptStepDetails() AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference
	// Experimental.
	DecryptStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsProperty
	// Experimental.
	DeleteStepDetails() AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsPropertyOutputReference
	// Experimental.
	DeleteStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TagStepDetails() AwsTransferWorkflow_OnExceptionStepsTagStepDetailsPropertyOutputReference
	// Experimental.
	TagStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsTagStepDetailsProperty
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
	PutCopyStepDetails(value *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty)
	// Experimental.
	PutCustomStepDetails(value *AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsProperty)
	// Experimental.
	PutDecryptStepDetails(value *AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsProperty)
	// Experimental.
	PutDeleteStepDetails(value *AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsProperty)
	// Experimental.
	PutTagStepDetails(value *AwsTransferWorkflow_OnExceptionStepsTagStepDetailsProperty)
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

// The jsii proxy struct for AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference
type jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) CopyStepDetails() AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference {
	var returns AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"copyStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) CopyStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty
	_jsii_.Get(
		j,
		"copyStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) CustomStepDetails() AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsPropertyOutputReference {
	var returns AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"customStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) CustomStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsProperty
	_jsii_.Get(
		j,
		"customStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) DecryptStepDetails() AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference {
	var returns AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"decryptStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) DecryptStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsProperty
	_jsii_.Get(
		j,
		"decryptStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) DeleteStepDetails() AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsPropertyOutputReference {
	var returns AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) DeleteStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsProperty
	_jsii_.Get(
		j,
		"deleteStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) TagStepDetails() AwsTransferWorkflow_OnExceptionStepsTagStepDetailsPropertyOutputReference {
	var returns AwsTransferWorkflow_OnExceptionStepsTagStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"tagStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) TagStepDetailsInput() *AwsTransferWorkflow_OnExceptionStepsTagStepDetailsProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsTagStepDetailsProperty
	_jsii_.Get(
		j,
		"tagStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTransferWorkflow_OnExceptionStepsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTransferWorkflow_OnExceptionStepsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferWorkflow.OnExceptionStepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTransferWorkflow_OnExceptionStepsPropertyOutputReference_Override(a AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferWorkflow.OnExceptionStepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) PutCopyStepDetails(value *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty) {
	if err := a.validatePutCopyStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCopyStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) PutCustomStepDetails(value *AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsProperty) {
	if err := a.validatePutCustomStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) PutDecryptStepDetails(value *AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsProperty) {
	if err := a.validatePutDecryptStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDecryptStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) PutDeleteStepDetails(value *AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsProperty) {
	if err := a.validatePutDeleteStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeleteStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) PutTagStepDetails(value *AwsTransferWorkflow_OnExceptionStepsTagStepDetailsProperty) {
	if err := a.validatePutTagStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ResetCopyStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ResetCustomStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ResetDecryptStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetDecryptStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ResetDeleteStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ResetTagStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetTagStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

