//go:build no_runtime_type_checking

package s3tables

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTable_MetadataPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTable_MetadataPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTable_MetadataPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTable_MetadataPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTable_MetadataPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTable_MetadataPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTable_MetadataPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTable_MetadataPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

