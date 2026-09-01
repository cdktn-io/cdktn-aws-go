package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbackup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbackup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBackupReportPlan_ReportSettingPropertyOutputReference interface {
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
	InternalValue() *AwsBackupReportPlan_ReportSettingProperty
	// Experimental.
	SetInternalValue(val *AwsBackupReportPlan_ReportSettingProperty)
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

// The jsii proxy struct for AwsBackupReportPlan_ReportSettingPropertyOutputReference
type jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) FrameworkArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frameworkArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) FrameworkArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frameworkArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) InternalValue() *AwsBackupReportPlan_ReportSettingProperty {
	var returns *AwsBackupReportPlan_ReportSettingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) NumberOfFrameworks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFrameworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) NumberOfFrameworksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFrameworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) OrganizationUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) OrganizationUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ReportTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ReportTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBackupReportPlan_ReportSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBackupReportPlan_ReportSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBackupReportPlan_ReportSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.AwsBackupReportPlan.ReportSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBackupReportPlan_ReportSettingPropertyOutputReference_Override(a AwsBackupReportPlan_ReportSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.AwsBackupReportPlan.ReportSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetFrameworkArns(val *[]*string) {
	if err := j.validateSetFrameworkArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameworkArns",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetInternalValue(val *AwsBackupReportPlan_ReportSettingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetNumberOfFrameworks(val *float64) {
	if err := j.validateSetNumberOfFrameworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfFrameworks",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetOrganizationUnits(val *[]*string) {
	if err := j.validateSetOrganizationUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationUnits",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetReportTemplate(val *string) {
	if err := j.validateSetReportTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportTemplate",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		a,
		"resetAccounts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ResetFrameworkArns() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameworkArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ResetNumberOfFrameworks() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberOfFrameworks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ResetOrganizationUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBackupReportPlan_ReportSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

