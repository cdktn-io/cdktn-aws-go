package kinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_OutputPropertyOutputReference interface {
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
	DestinationSchema() AwsApplication_DestinationSchemaPropertyOutputReference
	// Experimental.
	DestinationSchemaInput() *AwsApplication_DestinationSchemaProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KinesisFirehoseOutput() AwsApplication_KinesisFirehoseOutputPropertyOutputReference
	// Experimental.
	KinesisFirehoseOutputInput() *AwsApplication_KinesisFirehoseOutputProperty
	// Experimental.
	KinesisStreamsOutput() AwsApplication_KinesisStreamsOutputPropertyOutputReference
	// Experimental.
	KinesisStreamsOutputInput() *AwsApplication_KinesisStreamsOutputProperty
	// Experimental.
	LambdaOutput() AwsApplication_LambdaOutputPropertyOutputReference
	// Experimental.
	LambdaOutputInput() *AwsApplication_LambdaOutputProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	OutputId() *string
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
	PutDestinationSchema(value *AwsApplication_DestinationSchemaProperty)
	// Experimental.
	PutKinesisFirehoseOutput(value *AwsApplication_KinesisFirehoseOutputProperty)
	// Experimental.
	PutKinesisStreamsOutput(value *AwsApplication_KinesisStreamsOutputProperty)
	// Experimental.
	PutLambdaOutput(value *AwsApplication_LambdaOutputProperty)
	// Experimental.
	ResetKinesisFirehoseOutput()
	// Experimental.
	ResetKinesisStreamsOutput()
	// Experimental.
	ResetLambdaOutput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApplication_OutputPropertyOutputReference
type jsiiProxy_AwsApplication_OutputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) DestinationSchema() AwsApplication_DestinationSchemaPropertyOutputReference {
	var returns AwsApplication_DestinationSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"destinationSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) DestinationSchemaInput() *AwsApplication_DestinationSchemaProperty {
	var returns *AwsApplication_DestinationSchemaProperty
	_jsii_.Get(
		j,
		"destinationSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) KinesisFirehoseOutput() AwsApplication_KinesisFirehoseOutputPropertyOutputReference {
	var returns AwsApplication_KinesisFirehoseOutputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) KinesisFirehoseOutputInput() *AwsApplication_KinesisFirehoseOutputProperty {
	var returns *AwsApplication_KinesisFirehoseOutputProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) KinesisStreamsOutput() AwsApplication_KinesisStreamsOutputPropertyOutputReference {
	var returns AwsApplication_KinesisStreamsOutputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) KinesisStreamsOutputInput() *AwsApplication_KinesisStreamsOutputProperty {
	var returns *AwsApplication_KinesisStreamsOutputProperty
	_jsii_.Get(
		j,
		"kinesisStreamsOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) LambdaOutput() AwsApplication_LambdaOutputPropertyOutputReference {
	var returns AwsApplication_LambdaOutputPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) LambdaOutputInput() *AwsApplication_LambdaOutputProperty {
	var returns *AwsApplication_LambdaOutputProperty
	_jsii_.Get(
		j,
		"lambdaOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) OutputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_OutputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsApplication_OutputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_OutputPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_OutputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.OutputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_OutputPropertyOutputReference_Override(a AwsApplication_OutputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.OutputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_OutputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) PutDestinationSchema(value *AwsApplication_DestinationSchemaProperty) {
	if err := a.validatePutDestinationSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestinationSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) PutKinesisFirehoseOutput(value *AwsApplication_KinesisFirehoseOutputProperty) {
	if err := a.validatePutKinesisFirehoseOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehoseOutput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) PutKinesisStreamsOutput(value *AwsApplication_KinesisStreamsOutputProperty) {
	if err := a.validatePutKinesisStreamsOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamsOutput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) PutLambdaOutput(value *AwsApplication_LambdaOutputProperty) {
	if err := a.validatePutLambdaOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaOutput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ResetKinesisFirehoseOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehoseOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ResetKinesisStreamsOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamsOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ResetLambdaOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_OutputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

