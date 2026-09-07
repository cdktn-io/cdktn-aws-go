package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeliveryStream_ParquetSerDePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlockSizeBytes() *float64
	// Experimental.
	SetBlockSizeBytes(val *float64)
	// Experimental.
	BlockSizeBytesInput() *float64
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
	Compression() *string
	// Experimental.
	SetCompression(val *string)
	// Experimental.
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableDictionaryCompression() interface{}
	// Experimental.
	SetEnableDictionaryCompression(val interface{})
	// Experimental.
	EnableDictionaryCompressionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDeliveryStream_ParquetSerDeProperty
	// Experimental.
	SetInternalValue(val *AwsDeliveryStream_ParquetSerDeProperty)
	// Experimental.
	MaxPaddingBytes() *float64
	// Experimental.
	SetMaxPaddingBytes(val *float64)
	// Experimental.
	MaxPaddingBytesInput() *float64
	// Experimental.
	PageSizeBytes() *float64
	// Experimental.
	SetPageSizeBytes(val *float64)
	// Experimental.
	PageSizeBytesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WriterVersion() *string
	// Experimental.
	SetWriterVersion(val *string)
	// Experimental.
	WriterVersionInput() *string
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
	ResetBlockSizeBytes()
	// Experimental.
	ResetCompression()
	// Experimental.
	ResetEnableDictionaryCompression()
	// Experimental.
	ResetMaxPaddingBytes()
	// Experimental.
	ResetPageSizeBytes()
	// Experimental.
	ResetWriterVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeliveryStream_ParquetSerDePropertyOutputReference
type jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) BlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) BlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) EnableDictionaryCompression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDictionaryCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) EnableDictionaryCompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDictionaryCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) InternalValue() *AwsDeliveryStream_ParquetSerDeProperty {
	var returns *AwsDeliveryStream_ParquetSerDeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) MaxPaddingBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPaddingBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) MaxPaddingBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPaddingBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) PageSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) PageSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) WriterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) WriterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writerVersionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeliveryStream_ParquetSerDePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeliveryStream_ParquetSerDePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeliveryStream_ParquetSerDePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.ParquetSerDePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeliveryStream_ParquetSerDePropertyOutputReference_Override(a AwsDeliveryStream_ParquetSerDePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.ParquetSerDePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetBlockSizeBytes(val *float64) {
	if err := j.validateSetBlockSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetEnableDictionaryCompression(val interface{}) {
	if err := j.validateSetEnableDictionaryCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDictionaryCompression",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetInternalValue(val *AwsDeliveryStream_ParquetSerDeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetMaxPaddingBytes(val *float64) {
	if err := j.validateSetMaxPaddingBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPaddingBytes",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetPageSizeBytes(val *float64) {
	if err := j.validateSetPageSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pageSizeBytes",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference)SetWriterVersion(val *string) {
	if err := j.validateSetWriterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writerVersion",
		val,
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ResetBlockSizeBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockSizeBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		a,
		"resetCompression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ResetEnableDictionaryCompression() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDictionaryCompression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ResetMaxPaddingBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxPaddingBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ResetPageSizeBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetPageSizeBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ResetWriterVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetWriterVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_ParquetSerDePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

