// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AlertPolicyCompare = {
    GreaterThan: "GreaterThan",
    LessThan: "LessThan",
} as const;

export type AlertPolicyCompare = (typeof AlertPolicyCompare)[keyof typeof AlertPolicyCompare];

export const AlertPolicyType = {
    V1insightsdropletload1: "v1/insights/droplet/load_1",
    V1insightsdropletload5: "v1/insights/droplet/load_5",
    V1insightsdropletload15: "v1/insights/droplet/load_15",
    V1insightsdropletmemoryUtilizationPercent: "v1/insights/droplet/memory_utilization_percent",
    V1insightsdropletdiskUtilizationPercent: "v1/insights/droplet/disk_utilization_percent",
    V1insightsdropletcpu: "v1/insights/droplet/cpu",
    V1insightsdropletdiskRead: "v1/insights/droplet/disk_read",
    V1insightsdropletdiskWrite: "v1/insights/droplet/disk_write",
    V1insightsdropletpublicOutboundBandwidth: "v1/insights/droplet/public_outbound_bandwidth",
    V1insightsdropletpublicInboundBandwidth: "v1/insights/droplet/public_inbound_bandwidth",
    V1insightsdropletprivateOutboundBandwidth: "v1/insights/droplet/private_outbound_bandwidth",
    V1insightsdropletprivateInboundBandwidth: "v1/insights/droplet/private_inbound_bandwidth",
    V1insightslbaasavgCpuUtilizationPercent: "v1/insights/lbaas/avg_cpu_utilization_percent",
    V1insightslbaasconnectionUtilizationPercent: "v1/insights/lbaas/connection_utilization_percent",
    V1insightslbaasdropletHealth: "v1/insights/lbaas/droplet_health",
    V1insightslbaastlsConnectionsPerSecondUtilizationPercent: "v1/insights/lbaas/tls_connections_per_second_utilization_percent",
    V1insightslbaasincreaseInHttpErrorRatePercentage5xx: "v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx",
    V1insightslbaasincreaseInHttpErrorRatePercentage4xx: "v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx",
    V1insightslbaasincreaseInHttpErrorRateCount5xx: "v1/insights/lbaas/increase_in_http_error_rate_count_5xx",
    V1insightslbaasincreaseInHttpErrorRateCount4xx: "v1/insights/lbaas/increase_in_http_error_rate_count_4xx",
    V1insightslbaashighHttpRequestResponseTime: "v1/insights/lbaas/high_http_request_response_time",
    V1insightslbaashighHttpRequestResponseTime50p: "v1/insights/lbaas/high_http_request_response_time_50p",
    V1insightslbaashighHttpRequestResponseTime95p: "v1/insights/lbaas/high_http_request_response_time_95p",
    V1insightslbaashighHttpRequestResponseTime99p: "v1/insights/lbaas/high_http_request_response_time_99p",
    V1dbaasalertsload15Alerts: "v1/dbaas/alerts/load_15_alerts",
    V1dbaasalertsmemoryUtilizationAlerts: "v1/dbaas/alerts/memory_utilization_alerts",
    V1dbaasalertsdiskUtilizationAlerts: "v1/dbaas/alerts/disk_utilization_alerts",
    V1dbaasalertscpuAlerts: "v1/dbaas/alerts/cpu_alerts",
} as const;

export type AlertPolicyType = (typeof AlertPolicyType)[keyof typeof AlertPolicyType];

export const AlertPolicyWindow = {
    AlertPolicyWindow_5m: "5m",
    AlertPolicyWindow_10m: "10m",
    AlertPolicyWindow_30m: "30m",
    AlertPolicyWindow_1h: "1h",
} as const;

export type AlertPolicyWindow = (typeof AlertPolicyWindow)[keyof typeof AlertPolicyWindow];

export const Compare = {
    GreaterThan: "GreaterThan",
    LessThan: "LessThan",
} as const;

export type Compare = (typeof Compare)[keyof typeof Compare];

export const MetricsDataResultType = {
    Matrix: "matrix",
} as const;

export type MetricsDataResultType = (typeof MetricsDataResultType)[keyof typeof MetricsDataResultType];

export const MetricsStatus = {
    Success: "success",
    Error: "error",
} as const;

export type MetricsStatus = (typeof MetricsStatus)[keyof typeof MetricsStatus];

export const Type = {
    V1insightsdropletload1: "v1/insights/droplet/load_1",
    V1insightsdropletload5: "v1/insights/droplet/load_5",
    V1insightsdropletload15: "v1/insights/droplet/load_15",
    V1insightsdropletmemoryUtilizationPercent: "v1/insights/droplet/memory_utilization_percent",
    V1insightsdropletdiskUtilizationPercent: "v1/insights/droplet/disk_utilization_percent",
    V1insightsdropletcpu: "v1/insights/droplet/cpu",
    V1insightsdropletdiskRead: "v1/insights/droplet/disk_read",
    V1insightsdropletdiskWrite: "v1/insights/droplet/disk_write",
    V1insightsdropletpublicOutboundBandwidth: "v1/insights/droplet/public_outbound_bandwidth",
    V1insightsdropletpublicInboundBandwidth: "v1/insights/droplet/public_inbound_bandwidth",
    V1insightsdropletprivateOutboundBandwidth: "v1/insights/droplet/private_outbound_bandwidth",
    V1insightsdropletprivateInboundBandwidth: "v1/insights/droplet/private_inbound_bandwidth",
    V1insightslbaasavgCpuUtilizationPercent: "v1/insights/lbaas/avg_cpu_utilization_percent",
    V1insightslbaasconnectionUtilizationPercent: "v1/insights/lbaas/connection_utilization_percent",
    V1insightslbaasdropletHealth: "v1/insights/lbaas/droplet_health",
    V1insightslbaastlsConnectionsPerSecondUtilizationPercent: "v1/insights/lbaas/tls_connections_per_second_utilization_percent",
    V1insightslbaasincreaseInHttpErrorRatePercentage5xx: "v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx",
    V1insightslbaasincreaseInHttpErrorRatePercentage4xx: "v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx",
    V1insightslbaasincreaseInHttpErrorRateCount5xx: "v1/insights/lbaas/increase_in_http_error_rate_count_5xx",
    V1insightslbaasincreaseInHttpErrorRateCount4xx: "v1/insights/lbaas/increase_in_http_error_rate_count_4xx",
    V1insightslbaashighHttpRequestResponseTime: "v1/insights/lbaas/high_http_request_response_time",
    V1insightslbaashighHttpRequestResponseTime50p: "v1/insights/lbaas/high_http_request_response_time_50p",
    V1insightslbaashighHttpRequestResponseTime95p: "v1/insights/lbaas/high_http_request_response_time_95p",
    V1insightslbaashighHttpRequestResponseTime99p: "v1/insights/lbaas/high_http_request_response_time_99p",
    V1dbaasalertsload15Alerts: "v1/dbaas/alerts/load_15_alerts",
    V1dbaasalertsmemoryUtilizationAlerts: "v1/dbaas/alerts/memory_utilization_alerts",
    V1dbaasalertsdiskUtilizationAlerts: "v1/dbaas/alerts/disk_utilization_alerts",
    V1dbaasalertscpuAlerts: "v1/dbaas/alerts/cpu_alerts",
} as const;

export type Type = (typeof Type)[keyof typeof Type];

export const Window = {
    Window_5m: "5m",
    Window_10m: "10m",
    Window_30m: "30m",
    Window_1h: "1h",
} as const;

export type Window = (typeof Window)[keyof typeof Window];
