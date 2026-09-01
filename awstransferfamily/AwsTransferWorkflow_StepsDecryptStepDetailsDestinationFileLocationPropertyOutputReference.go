package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference interface {
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
	EfsFileLocation() AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
	// Experimental.
	EfsFileLocationInput() *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty
	// Experimental.
	SetInternalValue(val *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty)
	// Experimental.
	S3FileLocation() AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference
	// Experimental.
	S3FileLocationInput() *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty
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
	PutEfsFileLocation(value *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty)
	// Experimental.
	PutS3FileLocation(value *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty)
	// Experimental.
	ResetEfsFileLocation()
	// Experimental.
	ResetS3FileLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference
type jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) EfsFileLocation() AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference {
	var returns AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"efsFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) EfsFileLocationInput() *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty {
	var returns *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty
	_jsii_.Get(
		j,
		"efsFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) InternalValue() *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty {
	var returns *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) S3FileLocation() AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference {
	var returns AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3FileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) S3FileLocationInput() *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty {
	var returns *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty
	_jsii_.Get(
		j,
		"s3FileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferWorkflow.StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference_Override(a AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferWorkflow.StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetInternalValue(val *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) PutEfsFileLocation(value *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty) {
	if err := a.validatePutEfsFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEfsFileLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) PutS3FileLocation(value *AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty) {
	if err := a.validatePutS3FileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3FileLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ResetEfsFileLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetEfsFileLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ResetS3FileLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetS3FileLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTransferWorkflow_StepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

