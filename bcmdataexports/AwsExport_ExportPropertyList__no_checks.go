//go:build no_runtime_type_checking

package bcmdataexports

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsExport_ExportPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsExport_ExportPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsExport_ExportPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsExport_ExportPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsExport_ExportPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsExport_ExportPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsExport_ExportPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsExport_ExportPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

