package dynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTable_ImportTablePropertyOutputReference interface {
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
	InputCompressionType() *string
	// Experimental.
	SetInputCompressionType(val *string)
	// Experimental.
	InputCompressionTypeInput() *string
	// Experimental.
	InputFormat() *string
	// Experimental.
	SetInputFormat(val *string)
	// Experimental.
	InputFormatInput() *string
	// Experimental.
	InputFormatOptions() AwsTable_InputFormatOptionsPropertyOutputReference
	// Experimental.
	InputFormatOptionsInput() *AwsTable_InputFormatOptionsProperty
	// Experimental.
	InternalValue() *AwsTable_ImportTableProperty
	// Experimental.
	SetInternalValue(val *AwsTable_ImportTableProperty)
	// Experimental.
	S3BucketSource() AwsTable_S3BucketSourcePropertyOutputReference
	// Experimental.
	S3BucketSourceInput() *AwsTable_S3BucketSourceProperty
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
	PutInputFormatOptions(value *AwsTable_InputFormatOptionsProperty)
	// Experimental.
	PutS3BucketSource(value *AwsTable_S3BucketSourceProperty)
	// Experimental.
	ResetInputCompressionType()
	// Experimental.
	ResetInputFormatOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTable_ImportTablePropertyOutputReference
type jsiiProxy_AwsTable_ImportTablePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InputCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InputCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InputFormatOptions() AwsTable_InputFormatOptionsPropertyOutputReference {
	var returns AwsTable_InputFormatOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"inputFormatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InputFormatOptionsInput() *AwsTable_InputFormatOptionsProperty {
	var returns *AwsTable_InputFormatOptionsProperty
	_jsii_.Get(
		j,
		"inputFormatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InternalValue() *AwsTable_ImportTableProperty {
	var returns *AwsTable_ImportTableProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) S3BucketSource() AwsTable_S3BucketSourcePropertyOutputReference {
	var returns AwsTable_S3BucketSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"s3BucketSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) S3BucketSourceInput() *AwsTable_S3BucketSourceProperty {
	var returns *AwsTable_S3BucketSourceProperty
	_jsii_.Get(
		j,
		"s3BucketSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTable_ImportTablePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTable_ImportTablePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTable_ImportTablePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTable_ImportTablePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsTable.ImportTablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTable_ImportTablePropertyOutputReference_Override(a AwsTable_ImportTablePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsTable.ImportTablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetInputCompressionType(val *string) {
	if err := j.validateSetInputCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputCompressionType",
		val,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetInternalValue(val *AwsTable_ImportTableProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTable_ImportTablePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) PutInputFormatOptions(value *AwsTable_InputFormatOptionsProperty) {
	if err := a.validatePutInputFormatOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputFormatOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) PutS3BucketSource(value *AwsTable_S3BucketSourceProperty) {
	if err := a.validatePutS3BucketSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3BucketSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) ResetInputCompressionType() {
	_jsii_.InvokeVoid(
		a,
		"resetInputCompressionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) ResetInputFormatOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetInputFormatOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTable_ImportTablePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

