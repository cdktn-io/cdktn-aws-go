package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccountLevel() AwsStorageLensConfiguration_AccountLevelPropertyOutputReference
	// Experimental.
	AccountLevelInput() *AwsStorageLensConfiguration_AccountLevelProperty
	// Experimental.
	AwsOrg() AwsStorageLensConfiguration_AwsOrgPropertyOutputReference
	// Experimental.
	AwsOrgInput() *AwsStorageLensConfiguration_AwsOrgProperty
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
	DataExport() AwsStorageLensConfiguration_DataExportPropertyOutputReference
	// Experimental.
	DataExportInput() *AwsStorageLensConfiguration_DataExportProperty
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Exclude() AwsStorageLensConfiguration_ExcludePropertyOutputReference
	// Experimental.
	ExcludeInput() *AwsStorageLensConfiguration_ExcludeProperty
	// Experimental.
	ExpandedPrefixesDataExport() AwsStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
	// Experimental.
	ExpandedPrefixesDataExportInput() *AwsStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Include() AwsStorageLensConfiguration_IncludePropertyOutputReference
	// Experimental.
	IncludeInput() *AwsStorageLensConfiguration_IncludeProperty
	// Experimental.
	InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationProperty)
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
	PutAccountLevel(value *AwsStorageLensConfiguration_AccountLevelProperty)
	// Experimental.
	PutAwsOrg(value *AwsStorageLensConfiguration_AwsOrgProperty)
	// Experimental.
	PutDataExport(value *AwsStorageLensConfiguration_DataExportProperty)
	// Experimental.
	PutExclude(value *AwsStorageLensConfiguration_ExcludeProperty)
	// Experimental.
	PutExpandedPrefixesDataExport(value *AwsStorageLensConfiguration_ExpandedPrefixesDataExportProperty)
	// Experimental.
	PutInclude(value *AwsStorageLensConfiguration_IncludeProperty)
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

// The jsii proxy struct for AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AccountLevel() AwsStorageLensConfiguration_AccountLevelPropertyOutputReference {
	var returns AwsStorageLensConfiguration_AccountLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"accountLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AccountLevelInput() *AwsStorageLensConfiguration_AccountLevelProperty {
	var returns *AwsStorageLensConfiguration_AccountLevelProperty
	_jsii_.Get(
		j,
		"accountLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AwsOrg() AwsStorageLensConfiguration_AwsOrgPropertyOutputReference {
	var returns AwsStorageLensConfiguration_AwsOrgPropertyOutputReference
	_jsii_.Get(
		j,
		"awsOrg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AwsOrgInput() *AwsStorageLensConfiguration_AwsOrgProperty {
	var returns *AwsStorageLensConfiguration_AwsOrgProperty
	_jsii_.Get(
		j,
		"awsOrgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) DataExport() AwsStorageLensConfiguration_DataExportPropertyOutputReference {
	var returns AwsStorageLensConfiguration_DataExportPropertyOutputReference
	_jsii_.Get(
		j,
		"dataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) DataExportInput() *AwsStorageLensConfiguration_DataExportProperty {
	var returns *AwsStorageLensConfiguration_DataExportProperty
	_jsii_.Get(
		j,
		"dataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Exclude() AwsStorageLensConfiguration_ExcludePropertyOutputReference {
	var returns AwsStorageLensConfiguration_ExcludePropertyOutputReference
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExcludeInput() *AwsStorageLensConfiguration_ExcludeProperty {
	var returns *AwsStorageLensConfiguration_ExcludeProperty
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExpandedPrefixesDataExport() AwsStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference {
	var returns AwsStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
	_jsii_.Get(
		j,
		"expandedPrefixesDataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExpandedPrefixesDataExportInput() *AwsStorageLensConfiguration_ExpandedPrefixesDataExportProperty {
	var returns *AwsStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	_jsii_.Get(
		j,
		"expandedPrefixesDataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Include() AwsStorageLensConfiguration_IncludePropertyOutputReference {
	var returns AwsStorageLensConfiguration_IncludePropertyOutputReference
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) IncludeInput() *AwsStorageLensConfiguration_IncludeProperty {
	var returns *AwsStorageLensConfiguration_IncludeProperty
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PrefixDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PrefixDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference_Override(a AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetPrefixDelimiter(val *string) {
	if err := j.validateSetPrefixDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutAccountLevel(value *AwsStorageLensConfiguration_AccountLevelProperty) {
	if err := a.validatePutAccountLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutAwsOrg(value *AwsStorageLensConfiguration_AwsOrgProperty) {
	if err := a.validatePutAwsOrgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsOrg",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutDataExport(value *AwsStorageLensConfiguration_DataExportProperty) {
	if err := a.validatePutDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataExport",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutExclude(value *AwsStorageLensConfiguration_ExcludeProperty) {
	if err := a.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutExpandedPrefixesDataExport(value *AwsStorageLensConfiguration_ExpandedPrefixesDataExportProperty) {
	if err := a.validatePutExpandedPrefixesDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExpandedPrefixesDataExport",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutInclude(value *AwsStorageLensConfiguration_IncludeProperty) {
	if err := a.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetAwsOrg() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsOrg",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetDataExport() {
	_jsii_.InvokeVoid(
		a,
		"resetDataExport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		a,
		"resetExclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetExpandedPrefixesDataExport() {
	_jsii_.InvokeVoid(
		a,
		"resetExpandedPrefixesDataExport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetInclude() {
	_jsii_.InvokeVoid(
		a,
		"resetInclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetPrefixDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

