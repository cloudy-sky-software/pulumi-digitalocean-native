// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.DigitalOceanNative.MonitoringV2
{
    [EnumType]
    public readonly struct AlertPolicyCompare : IEquatable<AlertPolicyCompare>
    {
        private readonly string _value;

        private AlertPolicyCompare(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertPolicyCompare GreaterThan { get; } = new AlertPolicyCompare("GreaterThan");
        public static AlertPolicyCompare LessThan { get; } = new AlertPolicyCompare("LessThan");

        public static bool operator ==(AlertPolicyCompare left, AlertPolicyCompare right) => left.Equals(right);
        public static bool operator !=(AlertPolicyCompare left, AlertPolicyCompare right) => !left.Equals(right);

        public static explicit operator string(AlertPolicyCompare value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertPolicyCompare other && Equals(other);
        public bool Equals(AlertPolicyCompare other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct AlertPolicyType : IEquatable<AlertPolicyType>
    {
        private readonly string _value;

        private AlertPolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertPolicyType V1insightsdropletload1 { get; } = new AlertPolicyType("v1/insights/droplet/load_1");
        public static AlertPolicyType V1insightsdropletload5 { get; } = new AlertPolicyType("v1/insights/droplet/load_5");
        public static AlertPolicyType V1insightsdropletload15 { get; } = new AlertPolicyType("v1/insights/droplet/load_15");
        public static AlertPolicyType V1insightsdropletmemoryUtilizationPercent { get; } = new AlertPolicyType("v1/insights/droplet/memory_utilization_percent");
        public static AlertPolicyType V1insightsdropletdiskUtilizationPercent { get; } = new AlertPolicyType("v1/insights/droplet/disk_utilization_percent");
        public static AlertPolicyType V1insightsdropletcpu { get; } = new AlertPolicyType("v1/insights/droplet/cpu");
        public static AlertPolicyType V1insightsdropletdiskRead { get; } = new AlertPolicyType("v1/insights/droplet/disk_read");
        public static AlertPolicyType V1insightsdropletdiskWrite { get; } = new AlertPolicyType("v1/insights/droplet/disk_write");
        public static AlertPolicyType V1insightsdropletpublicOutboundBandwidth { get; } = new AlertPolicyType("v1/insights/droplet/public_outbound_bandwidth");
        public static AlertPolicyType V1insightsdropletpublicInboundBandwidth { get; } = new AlertPolicyType("v1/insights/droplet/public_inbound_bandwidth");
        public static AlertPolicyType V1insightsdropletprivateOutboundBandwidth { get; } = new AlertPolicyType("v1/insights/droplet/private_outbound_bandwidth");
        public static AlertPolicyType V1insightsdropletprivateInboundBandwidth { get; } = new AlertPolicyType("v1/insights/droplet/private_inbound_bandwidth");
        public static AlertPolicyType V1insightslbaasavgCpuUtilizationPercent { get; } = new AlertPolicyType("v1/insights/lbaas/avg_cpu_utilization_percent");
        public static AlertPolicyType V1insightslbaasconnectionUtilizationPercent { get; } = new AlertPolicyType("v1/insights/lbaas/connection_utilization_percent");
        public static AlertPolicyType V1insightslbaasdropletHealth { get; } = new AlertPolicyType("v1/insights/lbaas/droplet_health");
        public static AlertPolicyType V1insightslbaastlsConnectionsPerSecondUtilizationPercent { get; } = new AlertPolicyType("v1/insights/lbaas/tls_connections_per_second_utilization_percent");
        public static AlertPolicyType V1insightslbaasincreaseInHttpErrorRatePercentage5xx { get; } = new AlertPolicyType("v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx");
        public static AlertPolicyType V1insightslbaasincreaseInHttpErrorRatePercentage4xx { get; } = new AlertPolicyType("v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx");
        public static AlertPolicyType V1insightslbaasincreaseInHttpErrorRateCount5xx { get; } = new AlertPolicyType("v1/insights/lbaas/increase_in_http_error_rate_count_5xx");
        public static AlertPolicyType V1insightslbaasincreaseInHttpErrorRateCount4xx { get; } = new AlertPolicyType("v1/insights/lbaas/increase_in_http_error_rate_count_4xx");
        public static AlertPolicyType V1insightslbaashighHttpRequestResponseTime { get; } = new AlertPolicyType("v1/insights/lbaas/high_http_request_response_time");
        public static AlertPolicyType V1insightslbaashighHttpRequestResponseTime50p { get; } = new AlertPolicyType("v1/insights/lbaas/high_http_request_response_time_50p");
        public static AlertPolicyType V1insightslbaashighHttpRequestResponseTime95p { get; } = new AlertPolicyType("v1/insights/lbaas/high_http_request_response_time_95p");
        public static AlertPolicyType V1insightslbaashighHttpRequestResponseTime99p { get; } = new AlertPolicyType("v1/insights/lbaas/high_http_request_response_time_99p");
        public static AlertPolicyType V1dbaasalertsload15Alerts { get; } = new AlertPolicyType("v1/dbaas/alerts/load_15_alerts");
        public static AlertPolicyType V1dbaasalertsmemoryUtilizationAlerts { get; } = new AlertPolicyType("v1/dbaas/alerts/memory_utilization_alerts");
        public static AlertPolicyType V1dbaasalertsdiskUtilizationAlerts { get; } = new AlertPolicyType("v1/dbaas/alerts/disk_utilization_alerts");
        public static AlertPolicyType V1dbaasalertscpuAlerts { get; } = new AlertPolicyType("v1/dbaas/alerts/cpu_alerts");

        public static bool operator ==(AlertPolicyType left, AlertPolicyType right) => left.Equals(right);
        public static bool operator !=(AlertPolicyType left, AlertPolicyType right) => !left.Equals(right);

        public static explicit operator string(AlertPolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertPolicyType other && Equals(other);
        public bool Equals(AlertPolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct AlertPolicyWindow : IEquatable<AlertPolicyWindow>
    {
        private readonly string _value;

        private AlertPolicyWindow(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertPolicyWindow AlertPolicyWindow_5m { get; } = new AlertPolicyWindow("5m");
        public static AlertPolicyWindow AlertPolicyWindow_10m { get; } = new AlertPolicyWindow("10m");
        public static AlertPolicyWindow AlertPolicyWindow_30m { get; } = new AlertPolicyWindow("30m");
        public static AlertPolicyWindow AlertPolicyWindow_1h { get; } = new AlertPolicyWindow("1h");

        public static bool operator ==(AlertPolicyWindow left, AlertPolicyWindow right) => left.Equals(right);
        public static bool operator !=(AlertPolicyWindow left, AlertPolicyWindow right) => !left.Equals(right);

        public static explicit operator string(AlertPolicyWindow value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertPolicyWindow other && Equals(other);
        public bool Equals(AlertPolicyWindow other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Compare : IEquatable<Compare>
    {
        private readonly string _value;

        private Compare(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Compare GreaterThan { get; } = new Compare("GreaterThan");
        public static Compare LessThan { get; } = new Compare("LessThan");

        public static bool operator ==(Compare left, Compare right) => left.Equals(right);
        public static bool operator !=(Compare left, Compare right) => !left.Equals(right);

        public static explicit operator string(Compare value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Compare other && Equals(other);
        public bool Equals(Compare other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct MetricsDataResultType : IEquatable<MetricsDataResultType>
    {
        private readonly string _value;

        private MetricsDataResultType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MetricsDataResultType Matrix { get; } = new MetricsDataResultType("matrix");

        public static bool operator ==(MetricsDataResultType left, MetricsDataResultType right) => left.Equals(right);
        public static bool operator !=(MetricsDataResultType left, MetricsDataResultType right) => !left.Equals(right);

        public static explicit operator string(MetricsDataResultType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MetricsDataResultType other && Equals(other);
        public bool Equals(MetricsDataResultType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct MetricsStatus : IEquatable<MetricsStatus>
    {
        private readonly string _value;

        private MetricsStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MetricsStatus Success { get; } = new MetricsStatus("success");
        public static MetricsStatus Error { get; } = new MetricsStatus("error");

        public static bool operator ==(MetricsStatus left, MetricsStatus right) => left.Equals(right);
        public static bool operator !=(MetricsStatus left, MetricsStatus right) => !left.Equals(right);

        public static explicit operator string(MetricsStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MetricsStatus other && Equals(other);
        public bool Equals(MetricsStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Type : IEquatable<Type>
    {
        private readonly string _value;

        private Type(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Type V1insightsdropletload1 { get; } = new Type("v1/insights/droplet/load_1");
        public static Type V1insightsdropletload5 { get; } = new Type("v1/insights/droplet/load_5");
        public static Type V1insightsdropletload15 { get; } = new Type("v1/insights/droplet/load_15");
        public static Type V1insightsdropletmemoryUtilizationPercent { get; } = new Type("v1/insights/droplet/memory_utilization_percent");
        public static Type V1insightsdropletdiskUtilizationPercent { get; } = new Type("v1/insights/droplet/disk_utilization_percent");
        public static Type V1insightsdropletcpu { get; } = new Type("v1/insights/droplet/cpu");
        public static Type V1insightsdropletdiskRead { get; } = new Type("v1/insights/droplet/disk_read");
        public static Type V1insightsdropletdiskWrite { get; } = new Type("v1/insights/droplet/disk_write");
        public static Type V1insightsdropletpublicOutboundBandwidth { get; } = new Type("v1/insights/droplet/public_outbound_bandwidth");
        public static Type V1insightsdropletpublicInboundBandwidth { get; } = new Type("v1/insights/droplet/public_inbound_bandwidth");
        public static Type V1insightsdropletprivateOutboundBandwidth { get; } = new Type("v1/insights/droplet/private_outbound_bandwidth");
        public static Type V1insightsdropletprivateInboundBandwidth { get; } = new Type("v1/insights/droplet/private_inbound_bandwidth");
        public static Type V1insightslbaasavgCpuUtilizationPercent { get; } = new Type("v1/insights/lbaas/avg_cpu_utilization_percent");
        public static Type V1insightslbaasconnectionUtilizationPercent { get; } = new Type("v1/insights/lbaas/connection_utilization_percent");
        public static Type V1insightslbaasdropletHealth { get; } = new Type("v1/insights/lbaas/droplet_health");
        public static Type V1insightslbaastlsConnectionsPerSecondUtilizationPercent { get; } = new Type("v1/insights/lbaas/tls_connections_per_second_utilization_percent");
        public static Type V1insightslbaasincreaseInHttpErrorRatePercentage5xx { get; } = new Type("v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx");
        public static Type V1insightslbaasincreaseInHttpErrorRatePercentage4xx { get; } = new Type("v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx");
        public static Type V1insightslbaasincreaseInHttpErrorRateCount5xx { get; } = new Type("v1/insights/lbaas/increase_in_http_error_rate_count_5xx");
        public static Type V1insightslbaasincreaseInHttpErrorRateCount4xx { get; } = new Type("v1/insights/lbaas/increase_in_http_error_rate_count_4xx");
        public static Type V1insightslbaashighHttpRequestResponseTime { get; } = new Type("v1/insights/lbaas/high_http_request_response_time");
        public static Type V1insightslbaashighHttpRequestResponseTime50p { get; } = new Type("v1/insights/lbaas/high_http_request_response_time_50p");
        public static Type V1insightslbaashighHttpRequestResponseTime95p { get; } = new Type("v1/insights/lbaas/high_http_request_response_time_95p");
        public static Type V1insightslbaashighHttpRequestResponseTime99p { get; } = new Type("v1/insights/lbaas/high_http_request_response_time_99p");
        public static Type V1dbaasalertsload15Alerts { get; } = new Type("v1/dbaas/alerts/load_15_alerts");
        public static Type V1dbaasalertsmemoryUtilizationAlerts { get; } = new Type("v1/dbaas/alerts/memory_utilization_alerts");
        public static Type V1dbaasalertsdiskUtilizationAlerts { get; } = new Type("v1/dbaas/alerts/disk_utilization_alerts");
        public static Type V1dbaasalertscpuAlerts { get; } = new Type("v1/dbaas/alerts/cpu_alerts");

        public static bool operator ==(Type left, Type right) => left.Equals(right);
        public static bool operator !=(Type left, Type right) => !left.Equals(right);

        public static explicit operator string(Type value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Type other && Equals(other);
        public bool Equals(Type other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Window : IEquatable<Window>
    {
        private readonly string _value;

        private Window(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Window Window_5m { get; } = new Window("5m");
        public static Window Window_10m { get; } = new Window("10m");
        public static Window Window_30m { get; } = new Window("30m");
        public static Window Window_1h { get; } = new Window("1h");

        public static bool operator ==(Window left, Window right) => left.Equals(right);
        public static bool operator !=(Window left, Window right) => !left.Equals(right);

        public static explicit operator string(Window value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Window other && Equals(other);
        public bool Equals(Window other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
