//go:build no_runtime_type_checking

package s3control

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_RegionPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsMultiRegionAccessPoint_RegionPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

