package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference interface {
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
	DestinationFileLocation() AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationPropertyOutputReference
	// Experimental.
	DestinationFileLocationInput() *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty
	// Experimental.
	SetInternalValue(val *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty)
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
	PutDestinationFileLocation(value *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty)
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

// The jsii proxy struct for AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference
type jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) DestinationFileLocation() AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationPropertyOutputReference {
	var returns AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"destinationFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) DestinationFileLocationInput() *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty
	_jsii_.Get(
		j,
		"destinationFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) InternalValue() *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty {
	var returns *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) OverwriteExisting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteExisting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) OverwriteExistingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteExistingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) SourceFileLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) SourceFileLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferWorkflow.OnExceptionStepsCopyStepDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference_Override(a AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferWorkflow.OnExceptionStepsCopyStepDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetInternalValue(val *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetOverwriteExisting(val *string) {
	if err := j.validateSetOverwriteExistingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteExisting",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetSourceFileLocation(val *string) {
	if err := j.validateSetSourceFileLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFileLocation",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) PutDestinationFileLocation(value *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty) {
	if err := a.validatePutDestinationFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestinationFileLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ResetDestinationFileLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationFileLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ResetOverwriteExisting() {
	_jsii_.InvokeVoid(
		a,
		"resetOverwriteExisting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ResetSourceFileLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceFileLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

