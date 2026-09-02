package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfDbNodes_DbNodesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalDetails() *string
	// Experimental.
	Arn() *string
	// Experimental.
	BackupIpId() *string
	// Experimental.
	BackupVnic2Id() *string
	// Experimental.
	BackupVnicId() *string
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
	CpuCoreCount() *float64
	// Experimental.
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DbNodeStorageSize() *float64
	// Experimental.
	DbServerId() *string
	// Experimental.
	DbSystemId() *string
	// Experimental.
	FaultDomain() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HostIpId() *string
	// Experimental.
	Hostname() *string
	// Experimental.
	Id() *string
	// Experimental.
	InternalValue() *DataTfDbNodes_DbNodesProperty
	// Experimental.
	SetInternalValue(val *DataTfDbNodes_DbNodesProperty)
	// Experimental.
	MaintenanceType() *string
	// Experimental.
	MemorySize() *float64
	// Experimental.
	Ocid() *string
	// Experimental.
	OciResourceAnchorName() *string
	// Experimental.
	SoftwareStorageSize() *float64
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeMaintenanceWindowEnd() *string
	// Experimental.
	TimeMaintenanceWindowStart() *string
	// Experimental.
	TotalCpuCoreCount() *float64
	// Experimental.
	Vnic2Id() *string
	// Experimental.
	VnicId() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfDbNodes_DbNodesPropertyOutputReference
type jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) AdditionalDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) BackupIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) BackupVnic2Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVnic2Id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) BackupVnicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVnicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) DbNodeStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) DbServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) FaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) HostIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) InternalValue() *DataTfDbNodes_DbNodesProperty {
	var returns *DataTfDbNodes_DbNodesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) MaintenanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) SoftwareStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softwareStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) TimeMaintenanceWindowEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceWindowEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) TimeMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) TotalCpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Vnic2Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnic2Id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) VnicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnicId",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfDbNodes_DbNodesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfDbNodes_DbNodesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfDbNodes_DbNodesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataTfDbNodes.DbNodesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfDbNodes_DbNodesPropertyOutputReference_Override(d DataTfDbNodes_DbNodesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataTfDbNodes.DbNodesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference)SetInternalValue(val *DataTfDbNodes_DbNodesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDbNodes_DbNodesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

