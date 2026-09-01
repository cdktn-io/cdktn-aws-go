package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bitrate() *float64
	// Experimental.
	SetBitrate(val *float64)
	// Experimental.
	BitrateInput() *float64
	// Experimental.
	CodingMode() *string
	// Experimental.
	SetCodingMode(val *string)
	// Experimental.
	CodingModeInput() *string
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
	Dialnorm() *float64
	// Experimental.
	SetDialnorm(val *float64)
	// Experimental.
	DialnormInput() *float64
	// Experimental.
	DrcLine() *string
	// Experimental.
	SetDrcLine(val *string)
	// Experimental.
	DrcLineInput() *string
	// Experimental.
	DrcRf() *string
	// Experimental.
	SetDrcRf(val *string)
	// Experimental.
	DrcRfInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HeightTrim() *float64
	// Experimental.
	SetHeightTrim(val *float64)
	// Experimental.
	HeightTrimInput() *float64
	// Experimental.
	InternalValue() *AwsMedialiveChannel_Eac3AtmosSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_Eac3AtmosSettingsProperty)
	// Experimental.
	SurroundTrim() *float64
	// Experimental.
	SetSurroundTrim(val *float64)
	// Experimental.
	SurroundTrimInput() *float64
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
	ResetBitrate()
	// Experimental.
	ResetCodingMode()
	// Experimental.
	ResetDialnorm()
	// Experimental.
	ResetDrcLine()
	// Experimental.
	ResetDrcRf()
	// Experimental.
	ResetHeightTrim()
	// Experimental.
	ResetSurroundTrim()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) CodingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) CodingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) Dialnorm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnorm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) DialnormInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnormInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) DrcLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) DrcLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) DrcRf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) DrcRfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) HeightTrim() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightTrim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) HeightTrimInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightTrimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_Eac3AtmosSettingsProperty {
	var returns *AwsMedialiveChannel_Eac3AtmosSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) SurroundTrim() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"surroundTrim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) SurroundTrimInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"surroundTrimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.Eac3AtmosSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.Eac3AtmosSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetCodingMode(val *string) {
	if err := j.validateSetCodingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codingMode",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetDialnorm(val *float64) {
	if err := j.validateSetDialnormParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialnorm",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetDrcLine(val *string) {
	if err := j.validateSetDrcLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcLine",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetDrcRf(val *string) {
	if err := j.validateSetDrcRfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcRf",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetHeightTrim(val *float64) {
	if err := j.validateSetHeightTrimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heightTrim",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_Eac3AtmosSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetSurroundTrim(val *float64) {
	if err := j.validateSetSurroundTrimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surroundTrim",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		a,
		"resetBitrate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetCodingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetCodingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetDialnorm() {
	_jsii_.InvokeVoid(
		a,
		"resetDialnorm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetDrcLine() {
	_jsii_.InvokeVoid(
		a,
		"resetDrcLine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetDrcRf() {
	_jsii_.InvokeVoid(
		a,
		"resetDrcRf",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetHeightTrim() {
	_jsii_.InvokeVoid(
		a,
		"resetHeightTrim",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ResetSurroundTrim() {
	_jsii_.InvokeVoid(
		a,
		"resetSurroundTrim",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

