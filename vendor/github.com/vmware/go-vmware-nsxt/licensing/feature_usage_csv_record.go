/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package licensing

type FeatureUsageCsvRecord struct {

	// count of number of cpu sockets used by this feature
	CpuUsageCount int64 `json:"cpu_usage_count,omitempty"`

	// name of the feature
	Feature string `json:"feature,omitempty"`

	// count of number of vms used by this feature
	VmUsageCount int64 `json:"vm_usage_count,omitempty"`
}