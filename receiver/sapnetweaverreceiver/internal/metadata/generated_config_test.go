// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SapnetweaverAbapRfcCount:                MetricConfig{Enabled: true},
					SapnetweaverAbapSessionCount:            MetricConfig{Enabled: true},
					SapnetweaverAbapUpdateStatus:            MetricConfig{Enabled: true},
					SapnetweaverCacheEvictions:              MetricConfig{Enabled: true},
					SapnetweaverCacheHits:                   MetricConfig{Enabled: true},
					SapnetweaverCertificateValidity:         MetricConfig{Enabled: true},
					SapnetweaverConnectionErrorCount:        MetricConfig{Enabled: true},
					SapnetweaverCPUSystemUtilization:        MetricConfig{Enabled: true},
					SapnetweaverCPUUtilization:              MetricConfig{Enabled: true},
					SapnetweaverDatabaseDialogRequestTime:   MetricConfig{Enabled: true},
					SapnetweaverHostMemoryVirtualOverhead:   MetricConfig{Enabled: true},
					SapnetweaverHostMemoryVirtualSwap:       MetricConfig{Enabled: true},
					SapnetweaverHostSpoolListUtilization:    MetricConfig{Enabled: true},
					SapnetweaverLocksDequeueErrorsCount:     MetricConfig{Enabled: true},
					SapnetweaverLocksEnqueueCurrentCount:    MetricConfig{Enabled: true},
					SapnetweaverLocksEnqueueErrorsCount:     MetricConfig{Enabled: true},
					SapnetweaverLocksEnqueueHighCount:       MetricConfig{Enabled: true},
					SapnetweaverLocksEnqueueLockTime:        MetricConfig{Enabled: true},
					SapnetweaverLocksEnqueueLockWaitTime:    MetricConfig{Enabled: true},
					SapnetweaverLocksEnqueueMaxCount:        MetricConfig{Enabled: true},
					SapnetweaverMemoryConfigured:            MetricConfig{Enabled: true},
					SapnetweaverMemoryFree:                  MetricConfig{Enabled: true},
					SapnetweaverMemorySwapSpaceUtilization:  MetricConfig{Enabled: true},
					SapnetweaverProcessAvailability:         MetricConfig{Enabled: true},
					SapnetweaverQueueCount:                  MetricConfig{Enabled: true},
					SapnetweaverQueueMaxCount:               MetricConfig{Enabled: true},
					SapnetweaverQueuePeakCount:              MetricConfig{Enabled: true},
					SapnetweaverRequestCount:                MetricConfig{Enabled: true},
					SapnetweaverRequestTimeoutCount:         MetricConfig{Enabled: true},
					SapnetweaverResponseDuration:            MetricConfig{Enabled: true},
					SapnetweaverSessionCount:                MetricConfig{Enabled: true},
					SapnetweaverSessionsBrowserCount:        MetricConfig{Enabled: true},
					SapnetweaverSessionsEjbCount:            MetricConfig{Enabled: true},
					SapnetweaverSessionsHTTPCount:           MetricConfig{Enabled: true},
					SapnetweaverSessionsSecurityCount:       MetricConfig{Enabled: true},
					SapnetweaverSessionsWebCount:            MetricConfig{Enabled: true},
					SapnetweaverShortDumpsRate:              MetricConfig{Enabled: true},
					SapnetweaverSpoolRequestErrorCount:      MetricConfig{Enabled: true},
					SapnetweaverSystemInstanceAvailability:  MetricConfig{Enabled: true},
					SapnetweaverWorkProcessActiveCount:      MetricConfig{Enabled: true},
					SapnetweaverWorkProcessJobAbortedStatus: MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					SapnetweaverSID:      ResourceAttributeConfig{Enabled: true},
					SapnetweaverInstance: ResourceAttributeConfig{Enabled: true},
					SapnetweaverNode:     ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SapnetweaverAbapRfcCount:                MetricConfig{Enabled: false},
					SapnetweaverAbapSessionCount:            MetricConfig{Enabled: false},
					SapnetweaverAbapUpdateStatus:            MetricConfig{Enabled: false},
					SapnetweaverCacheEvictions:              MetricConfig{Enabled: false},
					SapnetweaverCacheHits:                   MetricConfig{Enabled: false},
					SapnetweaverCertificateValidity:         MetricConfig{Enabled: false},
					SapnetweaverConnectionErrorCount:        MetricConfig{Enabled: false},
					SapnetweaverCPUSystemUtilization:        MetricConfig{Enabled: false},
					SapnetweaverCPUUtilization:              MetricConfig{Enabled: false},
					SapnetweaverDatabaseDialogRequestTime:   MetricConfig{Enabled: false},
					SapnetweaverHostMemoryVirtualOverhead:   MetricConfig{Enabled: false},
					SapnetweaverHostMemoryVirtualSwap:       MetricConfig{Enabled: false},
					SapnetweaverHostSpoolListUtilization:    MetricConfig{Enabled: false},
					SapnetweaverLocksDequeueErrorsCount:     MetricConfig{Enabled: false},
					SapnetweaverLocksEnqueueCurrentCount:    MetricConfig{Enabled: false},
					SapnetweaverLocksEnqueueErrorsCount:     MetricConfig{Enabled: false},
					SapnetweaverLocksEnqueueHighCount:       MetricConfig{Enabled: false},
					SapnetweaverLocksEnqueueLockTime:        MetricConfig{Enabled: false},
					SapnetweaverLocksEnqueueLockWaitTime:    MetricConfig{Enabled: false},
					SapnetweaverLocksEnqueueMaxCount:        MetricConfig{Enabled: false},
					SapnetweaverMemoryConfigured:            MetricConfig{Enabled: false},
					SapnetweaverMemoryFree:                  MetricConfig{Enabled: false},
					SapnetweaverMemorySwapSpaceUtilization:  MetricConfig{Enabled: false},
					SapnetweaverProcessAvailability:         MetricConfig{Enabled: false},
					SapnetweaverQueueCount:                  MetricConfig{Enabled: false},
					SapnetweaverQueueMaxCount:               MetricConfig{Enabled: false},
					SapnetweaverQueuePeakCount:              MetricConfig{Enabled: false},
					SapnetweaverRequestCount:                MetricConfig{Enabled: false},
					SapnetweaverRequestTimeoutCount:         MetricConfig{Enabled: false},
					SapnetweaverResponseDuration:            MetricConfig{Enabled: false},
					SapnetweaverSessionCount:                MetricConfig{Enabled: false},
					SapnetweaverSessionsBrowserCount:        MetricConfig{Enabled: false},
					SapnetweaverSessionsEjbCount:            MetricConfig{Enabled: false},
					SapnetweaverSessionsHTTPCount:           MetricConfig{Enabled: false},
					SapnetweaverSessionsSecurityCount:       MetricConfig{Enabled: false},
					SapnetweaverSessionsWebCount:            MetricConfig{Enabled: false},
					SapnetweaverShortDumpsRate:              MetricConfig{Enabled: false},
					SapnetweaverSpoolRequestErrorCount:      MetricConfig{Enabled: false},
					SapnetweaverSystemInstanceAvailability:  MetricConfig{Enabled: false},
					SapnetweaverWorkProcessActiveCount:      MetricConfig{Enabled: false},
					SapnetweaverWorkProcessJobAbortedStatus: MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					SapnetweaverSID:      ResourceAttributeConfig{Enabled: false},
					SapnetweaverInstance: ResourceAttributeConfig{Enabled: false},
					SapnetweaverNode:     ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				SapnetweaverSID:      ResourceAttributeConfig{Enabled: true},
				SapnetweaverInstance: ResourceAttributeConfig{Enabled: true},
				SapnetweaverNode:     ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				SapnetweaverSID:      ResourceAttributeConfig{Enabled: false},
				SapnetweaverInstance: ResourceAttributeConfig{Enabled: false},
				SapnetweaverNode:     ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
