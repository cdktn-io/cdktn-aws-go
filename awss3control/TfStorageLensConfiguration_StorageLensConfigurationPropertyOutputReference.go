package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccountLevel() TfStorageLensConfiguration_AccountLevelPropertyOutputReference
	// Experimental.
	AccountLevelInput() *TfStorageLensConfiguration_AccountLevelProperty
	// Experimental.
	AwsOrg() TfStorageLensConfiguration_AwsOrgPropertyOutputReference
	// Experimental.
	AwsOrgInput() *TfStorageLensConfiguration_AwsOrgProperty
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
	DataExport() TfStorageLensConfiguration_DataExportPropertyOutputReference
	// Experimental.
	DataExportInput() *TfStorageLensConfiguration_DataExportProperty
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Exclude() TfStorageLensConfiguration_ExcludePropertyOutputReference
	// Experimental.
	ExcludeInput() *TfStorageLensConfiguration_ExcludeProperty
	// Experimental.
	ExpandedPrefixesDataExport() TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
	// Experimental.
	ExpandedPrefixesDataExportInput() *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Include() TfStorageLensConfiguration_IncludePropertyOutputReference
	// Experimental.
	IncludeInput() *TfStorageLensConfiguration_IncludeProperty
	// Experimental.
	InternalValue() *TfStorageLensConfiguration_StorageLensConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfStorageLensConfiguration_StorageLensConfigurationProperty)
	// Experimental.
	PrefixDelimiter() *string
	// Experimental.
	SetPrefixDelimiter(val *string)
	// Experimental.
	PrefixDelimiterInput() *string
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
	PutAccountLevel(value *TfStorageLensConfiguration_AccountLevelProperty)
	// Experimental.
	PutAwsOrg(value *TfStorageLensConfiguration_AwsOrgProperty)
	// Experimental.
	PutDataExport(value *TfStorageLensConfiguration_DataExportProperty)
	// Experimental.
	PutExclude(value *TfStorageLensConfiguration_ExcludeProperty)
	// Experimental.
	PutExpandedPrefixesDataExport(value *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty)
	// Experimental.
	PutInclude(value *TfStorageLensConfiguration_IncludeProperty)
	// Experimental.
	ResetAwsOrg()
	// Experimental.
	ResetDataExport()
	// Experimental.
	ResetExclude()
	// Experimental.
	ResetExpandedPrefixesDataExport()
	// Experimental.
	ResetInclude()
	// Experimental.
	ResetPrefixDelimiter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference
type jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AccountLevel() TfStorageLensConfiguration_AccountLevelPropertyOutputReference {
	var returns TfStorageLensConfiguration_AccountLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"accountLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AccountLevelInput() *TfStorageLensConfiguration_AccountLevelProperty {
	var returns *TfStorageLensConfiguration_AccountLevelProperty
	_jsii_.Get(
		j,
		"accountLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AwsOrg() TfStorageLensConfiguration_AwsOrgPropertyOutputReference {
	var returns TfStorageLensConfiguration_AwsOrgPropertyOutputReference
	_jsii_.Get(
		j,
		"awsOrg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AwsOrgInput() *TfStorageLensConfiguration_AwsOrgProperty {
	var returns *TfStorageLensConfiguration_AwsOrgProperty
	_jsii_.Get(
		j,
		"awsOrgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) DataExport() TfStorageLensConfiguration_DataExportPropertyOutputReference {
	var returns TfStorageLensConfiguration_DataExportPropertyOutputReference
	_jsii_.Get(
		j,
		"dataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) DataExportInput() *TfStorageLensConfiguration_DataExportProperty {
	var returns *TfStorageLensConfiguration_DataExportProperty
	_jsii_.Get(
		j,
		"dataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Exclude() TfStorageLensConfiguration_ExcludePropertyOutputReference {
	var returns TfStorageLensConfiguration_ExcludePropertyOutputReference
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExcludeInput() *TfStorageLensConfiguration_ExcludeProperty {
	var returns *TfStorageLensConfiguration_ExcludeProperty
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExpandedPrefixesDataExport() TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference {
	var returns TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
	_jsii_.Get(
		j,
		"expandedPrefixesDataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExpandedPrefixesDataExportInput() *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty {
	var returns *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	_jsii_.Get(
		j,
		"expandedPrefixesDataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Include() TfStorageLensConfiguration_IncludePropertyOutputReference {
	var returns TfStorageLensConfiguration_IncludePropertyOutputReference
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) IncludeInput() *TfStorageLensConfiguration_IncludeProperty {
	var returns *TfStorageLensConfiguration_IncludeProperty
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InternalValue() *TfStorageLensConfiguration_StorageLensConfigurationProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PrefixDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PrefixDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.StorageLensConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference_Override(t TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.StorageLensConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetInternalValue(val *TfStorageLensConfiguration_StorageLensConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetPrefixDelimiter(val *string) {
	if err := j.validateSetPrefixDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixDelimiter",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutAccountLevel(value *TfStorageLensConfiguration_AccountLevelProperty) {
	if err := t.validatePutAccountLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccountLevel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutAwsOrg(value *TfStorageLensConfiguration_AwsOrgProperty) {
	if err := t.validatePutAwsOrgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsOrg",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutDataExport(value *TfStorageLensConfiguration_DataExportProperty) {
	if err := t.validatePutDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataExport",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutExclude(value *TfStorageLensConfiguration_ExcludeProperty) {
	if err := t.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExclude",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutExpandedPrefixesDataExport(value *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty) {
	if err := t.validatePutExpandedPrefixesDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExpandedPrefixesDataExport",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutInclude(value *TfStorageLensConfiguration_IncludeProperty) {
	if err := t.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInclude",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetAwsOrg() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsOrg",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetDataExport() {
	_jsii_.InvokeVoid(
		t,
		"resetDataExport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		t,
		"resetExclude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetExpandedPrefixesDataExport() {
	_jsii_.InvokeVoid(
		t,
		"resetExpandedPrefixesDataExport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetInclude() {
	_jsii_.InvokeVoid(
		t,
		"resetInclude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetPrefixDelimiter() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefixDelimiter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

