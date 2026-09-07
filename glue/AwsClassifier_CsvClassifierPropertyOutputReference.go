package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsClassifier_CsvClassifierPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowSingleColumn() interface{}
	// Experimental.
	SetAllowSingleColumn(val interface{})
	// Experimental.
	AllowSingleColumnInput() interface{}
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
	ContainsHeader() *string
	// Experimental.
	SetContainsHeader(val *string)
	// Experimental.
	ContainsHeaderInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomDatatypeConfigured() interface{}
	// Experimental.
	SetCustomDatatypeConfigured(val interface{})
	// Experimental.
	CustomDatatypeConfiguredInput() interface{}
	// Experimental.
	CustomDatatypes() *[]*string
	// Experimental.
	SetCustomDatatypes(val *[]*string)
	// Experimental.
	CustomDatatypesInput() *[]*string
	// Experimental.
	Delimiter() *string
	// Experimental.
	SetDelimiter(val *string)
	// Experimental.
	DelimiterInput() *string
	// Experimental.
	DisableValueTrimming() interface{}
	// Experimental.
	SetDisableValueTrimming(val interface{})
	// Experimental.
	DisableValueTrimmingInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Header() *[]*string
	// Experimental.
	SetHeader(val *[]*string)
	// Experimental.
	HeaderInput() *[]*string
	// Experimental.
	InternalValue() *AwsClassifier_CsvClassifierProperty
	// Experimental.
	SetInternalValue(val *AwsClassifier_CsvClassifierProperty)
	// Experimental.
	QuoteSymbol() *string
	// Experimental.
	SetQuoteSymbol(val *string)
	// Experimental.
	QuoteSymbolInput() *string
	// Experimental.
	Serde() *string
	// Experimental.
	SetSerde(val *string)
	// Experimental.
	SerdeInput() *string
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
	ResetAllowSingleColumn()
	// Experimental.
	ResetContainsHeader()
	// Experimental.
	ResetCustomDatatypeConfigured()
	// Experimental.
	ResetCustomDatatypes()
	// Experimental.
	ResetDelimiter()
	// Experimental.
	ResetDisableValueTrimming()
	// Experimental.
	ResetHeader()
	// Experimental.
	ResetQuoteSymbol()
	// Experimental.
	ResetSerde()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsClassifier_CsvClassifierPropertyOutputReference
type jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) AllowSingleColumn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSingleColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) AllowSingleColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSingleColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ContainsHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ContainsHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) CustomDatatypeConfigured() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDatatypeConfigured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) CustomDatatypeConfiguredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDatatypeConfiguredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) CustomDatatypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDatatypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) CustomDatatypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDatatypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) DisableValueTrimming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableValueTrimming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) DisableValueTrimmingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableValueTrimmingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) Header() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) HeaderInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) InternalValue() *AwsClassifier_CsvClassifierProperty {
	var returns *AwsClassifier_CsvClassifierProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) QuoteSymbol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) QuoteSymbolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteSymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) Serde() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serde",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) SerdeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serdeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsClassifier_CsvClassifierPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsClassifier_CsvClassifierPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsClassifier_CsvClassifierPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsClassifier.CsvClassifierPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsClassifier_CsvClassifierPropertyOutputReference_Override(a AwsClassifier_CsvClassifierPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsClassifier.CsvClassifierPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetAllowSingleColumn(val interface{}) {
	if err := j.validateSetAllowSingleColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSingleColumn",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetContainsHeader(val *string) {
	if err := j.validateSetContainsHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containsHeader",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetCustomDatatypeConfigured(val interface{}) {
	if err := j.validateSetCustomDatatypeConfiguredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDatatypeConfigured",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetCustomDatatypes(val *[]*string) {
	if err := j.validateSetCustomDatatypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDatatypes",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetDisableValueTrimming(val interface{}) {
	if err := j.validateSetDisableValueTrimmingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableValueTrimming",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetHeader(val *[]*string) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetInternalValue(val *AwsClassifier_CsvClassifierProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetQuoteSymbol(val *string) {
	if err := j.validateSetQuoteSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quoteSymbol",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetSerde(val *string) {
	if err := j.validateSetSerdeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serde",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetAllowSingleColumn() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowSingleColumn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetContainsHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetContainsHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetCustomDatatypeConfigured() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDatatypeConfigured",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetCustomDatatypes() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDatatypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetDisableValueTrimming() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableValueTrimming",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetQuoteSymbol() {
	_jsii_.InvokeVoid(
		a,
		"resetQuoteSymbol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ResetSerde() {
	_jsii_.InvokeVoid(
		a,
		"resetSerde",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsClassifier_CsvClassifierPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

