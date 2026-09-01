package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlockSizeBytes() *float64
	// Experimental.
	SetBlockSizeBytes(val *float64)
	// Experimental.
	BlockSizeBytesInput() *float64
	// Experimental.
	BloomFilterColumns() *[]*string
	// Experimental.
	SetBloomFilterColumns(val *[]*string)
	// Experimental.
	BloomFilterColumnsInput() *[]*string
	// Experimental.
	BloomFilterFalsePositiveProbability() *float64
	// Experimental.
	SetBloomFilterFalsePositiveProbability(val *float64)
	// Experimental.
	BloomFilterFalsePositiveProbabilityInput() *float64
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
	DictionaryKeyThreshold() *float64
	// Experimental.
	SetDictionaryKeyThreshold(val *float64)
	// Experimental.
	DictionaryKeyThresholdInput() *float64
	// Experimental.
	EnablePadding() interface{}
	// Experimental.
	SetEnablePadding(val interface{})
	// Experimental.
	EnablePaddingInput() interface{}
	// Experimental.
	FormatVersion() *string
	// Experimental.
	SetFormatVersion(val *string)
	// Experimental.
	FormatVersionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsKinesisFirehoseDeliveryStream_OrcSerDeProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisFirehoseDeliveryStream_OrcSerDeProperty)
	// Experimental.
	PaddingTolerance() *float64
	// Experimental.
	SetPaddingTolerance(val *float64)
	// Experimental.
	PaddingToleranceInput() *float64
	// Experimental.
	RowIndexStride() *float64
	// Experimental.
	SetRowIndexStride(val *float64)
	// Experimental.
	RowIndexStrideInput() *float64
	// Experimental.
	StripeSizeBytes() *float64
	// Experimental.
	SetStripeSizeBytes(val *float64)
	// Experimental.
	StripeSizeBytesInput() *float64
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
	ResetBlockSizeBytes()
	// Experimental.
	ResetBloomFilterColumns()
	// Experimental.
	ResetBloomFilterFalsePositiveProbability()
	// Experimental.
	ResetCompression()
	// Experimental.
	ResetDictionaryKeyThreshold()
	// Experimental.
	ResetEnablePadding()
	// Experimental.
	ResetFormatVersion()
	// Experimental.
	ResetPaddingTolerance()
	// Experimental.
	ResetRowIndexStride()
	// Experimental.
	ResetStripeSizeBytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference
type jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) BlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) BlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) BloomFilterColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bloomFilterColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) BloomFilterColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bloomFilterColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) BloomFilterFalsePositiveProbability() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bloomFilterFalsePositiveProbability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) BloomFilterFalsePositiveProbabilityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bloomFilterFalsePositiveProbabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) DictionaryKeyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictionaryKeyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) DictionaryKeyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictionaryKeyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) EnablePadding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) EnablePaddingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) FormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) FormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) InternalValue() *AwsKinesisFirehoseDeliveryStream_OrcSerDeProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_OrcSerDeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) PaddingTolerance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"paddingTolerance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) PaddingToleranceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"paddingToleranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) RowIndexStride() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowIndexStride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) RowIndexStrideInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowIndexStrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) StripeSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stripeSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) StripeSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stripeSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.OrcSerDePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference_Override(a AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.OrcSerDePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetBlockSizeBytes(val *float64) {
	if err := j.validateSetBlockSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetBloomFilterColumns(val *[]*string) {
	if err := j.validateSetBloomFilterColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bloomFilterColumns",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetBloomFilterFalsePositiveProbability(val *float64) {
	if err := j.validateSetBloomFilterFalsePositiveProbabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bloomFilterFalsePositiveProbability",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetDictionaryKeyThreshold(val *float64) {
	if err := j.validateSetDictionaryKeyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dictionaryKeyThreshold",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetEnablePadding(val interface{}) {
	if err := j.validateSetEnablePaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePadding",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetFormatVersion(val *string) {
	if err := j.validateSetFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatVersion",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetInternalValue(val *AwsKinesisFirehoseDeliveryStream_OrcSerDeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetPaddingTolerance(val *float64) {
	if err := j.validateSetPaddingToleranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paddingTolerance",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetRowIndexStride(val *float64) {
	if err := j.validateSetRowIndexStrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowIndexStride",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetStripeSizeBytes(val *float64) {
	if err := j.validateSetStripeSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripeSizeBytes",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetBlockSizeBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockSizeBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetBloomFilterColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetBloomFilterColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetBloomFilterFalsePositiveProbability() {
	_jsii_.InvokeVoid(
		a,
		"resetBloomFilterFalsePositiveProbability",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		a,
		"resetCompression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetDictionaryKeyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetDictionaryKeyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetEnablePadding() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePadding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetFormatVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetFormatVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetPaddingTolerance() {
	_jsii_.InvokeVoid(
		a,
		"resetPaddingTolerance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetRowIndexStride() {
	_jsii_.InvokeVoid(
		a,
		"resetRowIndexStride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ResetStripeSizeBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetStripeSizeBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_OrcSerDePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

