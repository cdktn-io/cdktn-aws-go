package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbackup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbackup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReportPlan_ReportSettingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Accounts() *[]*string
	// Experimental.
	SetAccounts(val *[]*string)
	// Experimental.
	AccountsInput() *[]*string
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
	FrameworkArns() *[]*string
	// Experimental.
	SetFrameworkArns(val *[]*string)
	// Experimental.
	FrameworkArnsInput() *[]*string
	// Experimental.
	InternalValue() *TfReportPlan_ReportSettingProperty
	// Experimental.
	SetInternalValue(val *TfReportPlan_ReportSettingProperty)
	// Experimental.
	NumberOfFrameworks() *float64
	// Experimental.
	SetNumberOfFrameworks(val *float64)
	// Experimental.
	NumberOfFrameworksInput() *float64
	// Experimental.
	OrganizationUnits() *[]*string
	// Experimental.
	SetOrganizationUnits(val *[]*string)
	// Experimental.
	OrganizationUnitsInput() *[]*string
	// Experimental.
	Regions() *[]*string
	// Experimental.
	SetRegions(val *[]*string)
	// Experimental.
	RegionsInput() *[]*string
	// Experimental.
	ReportTemplate() *string
	// Experimental.
	SetReportTemplate(val *string)
	// Experimental.
	ReportTemplateInput() *string
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
	ResetAccounts()
	// Experimental.
	ResetFrameworkArns()
	// Experimental.
	ResetNumberOfFrameworks()
	// Experimental.
	ResetOrganizationUnits()
	// Experimental.
	ResetRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfReportPlan_ReportSettingPropertyOutputReference
type jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) FrameworkArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frameworkArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) FrameworkArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frameworkArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) InternalValue() *TfReportPlan_ReportSettingProperty {
	var returns *TfReportPlan_ReportSettingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) NumberOfFrameworks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFrameworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) NumberOfFrameworksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFrameworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) OrganizationUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) OrganizationUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ReportTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ReportTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReportPlan_ReportSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfReportPlan_ReportSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReportPlan_ReportSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.TfReportPlan.ReportSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReportPlan_ReportSettingPropertyOutputReference_Override(t TfReportPlan_ReportSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.TfReportPlan.ReportSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetFrameworkArns(val *[]*string) {
	if err := j.validateSetFrameworkArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameworkArns",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetInternalValue(val *TfReportPlan_ReportSettingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetNumberOfFrameworks(val *float64) {
	if err := j.validateSetNumberOfFrameworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfFrameworks",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetOrganizationUnits(val *[]*string) {
	if err := j.validateSetOrganizationUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationUnits",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetReportTemplate(val *string) {
	if err := j.validateSetReportTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportTemplate",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		t,
		"resetAccounts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ResetFrameworkArns() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameworkArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ResetNumberOfFrameworks() {
	_jsii_.InvokeVoid(
		t,
		"resetNumberOfFrameworks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ResetOrganizationUnits() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationUnits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		t,
		"resetRegions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReportPlan_ReportSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

