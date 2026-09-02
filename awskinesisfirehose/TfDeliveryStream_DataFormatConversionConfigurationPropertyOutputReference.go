package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference interface {
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InputFormatConfiguration() TfDeliveryStream_InputFormatConfigurationPropertyOutputReference
	// Experimental.
	InputFormatConfigurationInput() *TfDeliveryStream_InputFormatConfigurationProperty
	// Experimental.
	InternalValue() *TfDeliveryStream_DataFormatConversionConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_DataFormatConversionConfigurationProperty)
	// Experimental.
	OutputFormatConfiguration() TfDeliveryStream_OutputFormatConfigurationPropertyOutputReference
	// Experimental.
	OutputFormatConfigurationInput() *TfDeliveryStream_OutputFormatConfigurationProperty
	// Experimental.
	SchemaConfiguration() TfDeliveryStream_SchemaConfigurationPropertyOutputReference
	// Experimental.
	SchemaConfigurationInput() *TfDeliveryStream_SchemaConfigurationProperty
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
	PutInputFormatConfiguration(value *TfDeliveryStream_InputFormatConfigurationProperty)
	// Experimental.
	PutOutputFormatConfiguration(value *TfDeliveryStream_OutputFormatConfigurationProperty)
	// Experimental.
	PutSchemaConfiguration(value *TfDeliveryStream_SchemaConfigurationProperty)
	// Experimental.
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference
type jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) InputFormatConfiguration() TfDeliveryStream_InputFormatConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_InputFormatConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"inputFormatConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) InputFormatConfigurationInput() *TfDeliveryStream_InputFormatConfigurationProperty {
	var returns *TfDeliveryStream_InputFormatConfigurationProperty
	_jsii_.Get(
		j,
		"inputFormatConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) InternalValue() *TfDeliveryStream_DataFormatConversionConfigurationProperty {
	var returns *TfDeliveryStream_DataFormatConversionConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) OutputFormatConfiguration() TfDeliveryStream_OutputFormatConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_OutputFormatConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"outputFormatConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) OutputFormatConfigurationInput() *TfDeliveryStream_OutputFormatConfigurationProperty {
	var returns *TfDeliveryStream_OutputFormatConfigurationProperty
	_jsii_.Get(
		j,
		"outputFormatConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) SchemaConfiguration() TfDeliveryStream_SchemaConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SchemaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) SchemaConfigurationInput() *TfDeliveryStream_SchemaConfigurationProperty {
	var returns *TfDeliveryStream_SchemaConfigurationProperty
	_jsii_.Get(
		j,
		"schemaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.DataFormatConversionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference_Override(t TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.DataFormatConversionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_DataFormatConversionConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) PutInputFormatConfiguration(value *TfDeliveryStream_InputFormatConfigurationProperty) {
	if err := t.validatePutInputFormatConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputFormatConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) PutOutputFormatConfiguration(value *TfDeliveryStream_OutputFormatConfigurationProperty) {
	if err := t.validatePutOutputFormatConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputFormatConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) PutSchemaConfiguration(value *TfDeliveryStream_SchemaConfigurationProperty) {
	if err := t.validatePutSchemaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchemaConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

