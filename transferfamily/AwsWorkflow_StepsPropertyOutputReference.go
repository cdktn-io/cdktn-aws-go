package transferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/transferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/transferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkflow_StepsPropertyOutputReference interface {
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
	CopyStepDetails() AwsWorkflow_StepsCopyStepDetailsPropertyOutputReference
	// Experimental.
	CopyStepDetailsInput() *AwsWorkflow_StepsCopyStepDetailsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomStepDetails() AwsWorkflow_StepsCustomStepDetailsPropertyOutputReference
	// Experimental.
	CustomStepDetailsInput() *AwsWorkflow_StepsCustomStepDetailsProperty
	// Experimental.
	DecryptStepDetails() AwsWorkflow_StepsDecryptStepDetailsPropertyOutputReference
	// Experimental.
	DecryptStepDetailsInput() *AwsWorkflow_StepsDecryptStepDetailsProperty
	// Experimental.
	DeleteStepDetails() AwsWorkflow_StepsDeleteStepDetailsPropertyOutputReference
	// Experimental.
	DeleteStepDetailsInput() *AwsWorkflow_StepsDeleteStepDetailsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TagStepDetails() AwsWorkflow_StepsTagStepDetailsPropertyOutputReference
	// Experimental.
	TagStepDetailsInput() *AwsWorkflow_StepsTagStepDetailsProperty
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
	PutCopyStepDetails(value *AwsWorkflow_StepsCopyStepDetailsProperty)
	// Experimental.
	PutCustomStepDetails(value *AwsWorkflow_StepsCustomStepDetailsProperty)
	// Experimental.
	PutDecryptStepDetails(value *AwsWorkflow_StepsDecryptStepDetailsProperty)
	// Experimental.
	PutDeleteStepDetails(value *AwsWorkflow_StepsDeleteStepDetailsProperty)
	// Experimental.
	PutTagStepDetails(value *AwsWorkflow_StepsTagStepDetailsProperty)
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

// The jsii proxy struct for AwsWorkflow_StepsPropertyOutputReference
type jsiiProxy_AwsWorkflow_StepsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) CopyStepDetails() AwsWorkflow_StepsCopyStepDetailsPropertyOutputReference {
	var returns AwsWorkflow_StepsCopyStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"copyStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) CopyStepDetailsInput() *AwsWorkflow_StepsCopyStepDetailsProperty {
	var returns *AwsWorkflow_StepsCopyStepDetailsProperty
	_jsii_.Get(
		j,
		"copyStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) CustomStepDetails() AwsWorkflow_StepsCustomStepDetailsPropertyOutputReference {
	var returns AwsWorkflow_StepsCustomStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"customStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) CustomStepDetailsInput() *AwsWorkflow_StepsCustomStepDetailsProperty {
	var returns *AwsWorkflow_StepsCustomStepDetailsProperty
	_jsii_.Get(
		j,
		"customStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) DecryptStepDetails() AwsWorkflow_StepsDecryptStepDetailsPropertyOutputReference {
	var returns AwsWorkflow_StepsDecryptStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"decryptStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) DecryptStepDetailsInput() *AwsWorkflow_StepsDecryptStepDetailsProperty {
	var returns *AwsWorkflow_StepsDecryptStepDetailsProperty
	_jsii_.Get(
		j,
		"decryptStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) DeleteStepDetails() AwsWorkflow_StepsDeleteStepDetailsPropertyOutputReference {
	var returns AwsWorkflow_StepsDeleteStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"deleteStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) DeleteStepDetailsInput() *AwsWorkflow_StepsDeleteStepDetailsProperty {
	var returns *AwsWorkflow_StepsDeleteStepDetailsProperty
	_jsii_.Get(
		j,
		"deleteStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) TagStepDetails() AwsWorkflow_StepsTagStepDetailsPropertyOutputReference {
	var returns AwsWorkflow_StepsTagStepDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"tagStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) TagStepDetailsInput() *AwsWorkflow_StepsTagStepDetailsProperty {
	var returns *AwsWorkflow_StepsTagStepDetailsProperty
	_jsii_.Get(
		j,
		"tagStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWorkflow_StepsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWorkflow_StepsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWorkflow_StepsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkflow_StepsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsWorkflow.StepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWorkflow_StepsPropertyOutputReference_Override(a AwsWorkflow_StepsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsWorkflow.StepsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) PutCopyStepDetails(value *AwsWorkflow_StepsCopyStepDetailsProperty) {
	if err := a.validatePutCopyStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCopyStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) PutCustomStepDetails(value *AwsWorkflow_StepsCustomStepDetailsProperty) {
	if err := a.validatePutCustomStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) PutDecryptStepDetails(value *AwsWorkflow_StepsDecryptStepDetailsProperty) {
	if err := a.validatePutDecryptStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDecryptStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) PutDeleteStepDetails(value *AwsWorkflow_StepsDeleteStepDetailsProperty) {
	if err := a.validatePutDeleteStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeleteStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) PutTagStepDetails(value *AwsWorkflow_StepsTagStepDetailsProperty) {
	if err := a.validatePutTagStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagStepDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ResetCopyStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ResetCustomStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ResetDecryptStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetDecryptStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ResetDeleteStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ResetTagStepDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetTagStepDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWorkflow_StepsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

