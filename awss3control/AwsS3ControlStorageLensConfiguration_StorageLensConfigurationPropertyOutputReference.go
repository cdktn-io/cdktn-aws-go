package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccountLevel() AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference
	// Experimental.
	AccountLevelInput() *AwsS3ControlStorageLensConfiguration_AccountLevelProperty
	// Experimental.
	AwsOrg() AwsS3ControlStorageLensConfiguration_AwsOrgPropertyOutputReference
	// Experimental.
	AwsOrgInput() *AwsS3ControlStorageLensConfiguration_AwsOrgProperty
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
	DataExport() AwsS3ControlStorageLensConfiguration_DataExportPropertyOutputReference
	// Experimental.
	DataExportInput() *AwsS3ControlStorageLensConfiguration_DataExportProperty
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Exclude() AwsS3ControlStorageLensConfiguration_ExcludePropertyOutputReference
	// Experimental.
	ExcludeInput() *AwsS3ControlStorageLensConfiguration_ExcludeProperty
	// Experimental.
	ExpandedPrefixesDataExport() AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
	// Experimental.
	ExpandedPrefixesDataExportInput() *AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Include() AwsS3ControlStorageLensConfiguration_IncludePropertyOutputReference
	// Experimental.
	IncludeInput() *AwsS3ControlStorageLensConfiguration_IncludeProperty
	// Experimental.
	InternalValue() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationProperty)
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
	PutAccountLevel(value *AwsS3ControlStorageLensConfiguration_AccountLevelProperty)
	// Experimental.
	PutAwsOrg(value *AwsS3ControlStorageLensConfiguration_AwsOrgProperty)
	// Experimental.
	PutDataExport(value *AwsS3ControlStorageLensConfiguration_DataExportProperty)
	// Experimental.
	PutExclude(value *AwsS3ControlStorageLensConfiguration_ExcludeProperty)
	// Experimental.
	PutExpandedPrefixesDataExport(value *AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty)
	// Experimental.
	PutInclude(value *AwsS3ControlStorageLensConfiguration_IncludeProperty)
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

// The jsii proxy struct for AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference
type jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AccountLevel() AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"accountLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AccountLevelInput() *AwsS3ControlStorageLensConfiguration_AccountLevelProperty {
	var returns *AwsS3ControlStorageLensConfiguration_AccountLevelProperty
	_jsii_.Get(
		j,
		"accountLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AwsOrg() AwsS3ControlStorageLensConfiguration_AwsOrgPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_AwsOrgPropertyOutputReference
	_jsii_.Get(
		j,
		"awsOrg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) AwsOrgInput() *AwsS3ControlStorageLensConfiguration_AwsOrgProperty {
	var returns *AwsS3ControlStorageLensConfiguration_AwsOrgProperty
	_jsii_.Get(
		j,
		"awsOrgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) DataExport() AwsS3ControlStorageLensConfiguration_DataExportPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_DataExportPropertyOutputReference
	_jsii_.Get(
		j,
		"dataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) DataExportInput() *AwsS3ControlStorageLensConfiguration_DataExportProperty {
	var returns *AwsS3ControlStorageLensConfiguration_DataExportProperty
	_jsii_.Get(
		j,
		"dataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Exclude() AwsS3ControlStorageLensConfiguration_ExcludePropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_ExcludePropertyOutputReference
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExcludeInput() *AwsS3ControlStorageLensConfiguration_ExcludeProperty {
	var returns *AwsS3ControlStorageLensConfiguration_ExcludeProperty
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExpandedPrefixesDataExport() AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
	_jsii_.Get(
		j,
		"expandedPrefixesDataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ExpandedPrefixesDataExportInput() *AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty {
	var returns *AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	_jsii_.Get(
		j,
		"expandedPrefixesDataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Include() AwsS3ControlStorageLensConfiguration_IncludePropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_IncludePropertyOutputReference
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) IncludeInput() *AwsS3ControlStorageLensConfiguration_IncludeProperty {
	var returns *AwsS3ControlStorageLensConfiguration_IncludeProperty
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InternalValue() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PrefixDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PrefixDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.StorageLensConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference_Override(a AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.StorageLensConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetInternalValue(val *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetPrefixDelimiter(val *string) {
	if err := j.validateSetPrefixDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutAccountLevel(value *AwsS3ControlStorageLensConfiguration_AccountLevelProperty) {
	if err := a.validatePutAccountLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutAwsOrg(value *AwsS3ControlStorageLensConfiguration_AwsOrgProperty) {
	if err := a.validatePutAwsOrgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsOrg",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutDataExport(value *AwsS3ControlStorageLensConfiguration_DataExportProperty) {
	if err := a.validatePutDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataExport",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutExclude(value *AwsS3ControlStorageLensConfiguration_ExcludeProperty) {
	if err := a.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutExpandedPrefixesDataExport(value *AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty) {
	if err := a.validatePutExpandedPrefixesDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExpandedPrefixesDataExport",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) PutInclude(value *AwsS3ControlStorageLensConfiguration_IncludeProperty) {
	if err := a.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetAwsOrg() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsOrg",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetDataExport() {
	_jsii_.InvokeVoid(
		a,
		"resetDataExport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		a,
		"resetExclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetExpandedPrefixesDataExport() {
	_jsii_.InvokeVoid(
		a,
		"resetExpandedPrefixesDataExport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetInclude() {
	_jsii_.InvokeVoid(
		a,
		"resetInclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ResetPrefixDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

