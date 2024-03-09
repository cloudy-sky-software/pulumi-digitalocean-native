// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.DigitalOceanNative.LoadBalancersV2
{
    /// <summary>
    /// The protocol used for traffic to the load balancer. The possible values are: `http`, `https`, `http2`, `http3`, `tcp`, or `udp`. If you set the  `entry_protocol` to `udp`, the `target_protocol` must be set to `udp`.  When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
    /// </summary>
    [EnumType]
    public readonly struct ForwardingRuleEntryProtocol : IEquatable<ForwardingRuleEntryProtocol>
    {
        private readonly string _value;

        private ForwardingRuleEntryProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ForwardingRuleEntryProtocol Http { get; } = new ForwardingRuleEntryProtocol("http");
        public static ForwardingRuleEntryProtocol Https { get; } = new ForwardingRuleEntryProtocol("https");
        public static ForwardingRuleEntryProtocol Http2 { get; } = new ForwardingRuleEntryProtocol("http2");
        public static ForwardingRuleEntryProtocol Http3 { get; } = new ForwardingRuleEntryProtocol("http3");
        public static ForwardingRuleEntryProtocol Tcp { get; } = new ForwardingRuleEntryProtocol("tcp");
        public static ForwardingRuleEntryProtocol Udp { get; } = new ForwardingRuleEntryProtocol("udp");

        public static bool operator ==(ForwardingRuleEntryProtocol left, ForwardingRuleEntryProtocol right) => left.Equals(right);
        public static bool operator !=(ForwardingRuleEntryProtocol left, ForwardingRuleEntryProtocol right) => !left.Equals(right);

        public static explicit operator string(ForwardingRuleEntryProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ForwardingRuleEntryProtocol other && Equals(other);
        public bool Equals(ForwardingRuleEntryProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The protocol used for traffic from the load balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, `tcp`, or `udp`. If you set the `target_protocol` to `udp`, the `entry_protocol` must be set to  `udp`. When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
    /// </summary>
    [EnumType]
    public readonly struct ForwardingRuleTargetProtocol : IEquatable<ForwardingRuleTargetProtocol>
    {
        private readonly string _value;

        private ForwardingRuleTargetProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ForwardingRuleTargetProtocol Http { get; } = new ForwardingRuleTargetProtocol("http");
        public static ForwardingRuleTargetProtocol Https { get; } = new ForwardingRuleTargetProtocol("https");
        public static ForwardingRuleTargetProtocol Http2 { get; } = new ForwardingRuleTargetProtocol("http2");
        public static ForwardingRuleTargetProtocol Tcp { get; } = new ForwardingRuleTargetProtocol("tcp");
        public static ForwardingRuleTargetProtocol Udp { get; } = new ForwardingRuleTargetProtocol("udp");

        public static bool operator ==(ForwardingRuleTargetProtocol left, ForwardingRuleTargetProtocol right) => left.Equals(right);
        public static bool operator !=(ForwardingRuleTargetProtocol left, ForwardingRuleTargetProtocol right) => !left.Equals(right);

        public static explicit operator string(ForwardingRuleTargetProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ForwardingRuleTargetProtocol other && Equals(other);
        public bool Equals(ForwardingRuleTargetProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The protocol used for health checks sent to the backend Droplets. The possible values are `http`, `https`, or `tcp`.
    /// </summary>
    [EnumType]
    public readonly struct HealthCheckProtocol : IEquatable<HealthCheckProtocol>
    {
        private readonly string _value;

        private HealthCheckProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HealthCheckProtocol Http { get; } = new HealthCheckProtocol("http");
        public static HealthCheckProtocol Https { get; } = new HealthCheckProtocol("https");
        public static HealthCheckProtocol Tcp { get; } = new HealthCheckProtocol("tcp");

        public static bool operator ==(HealthCheckProtocol left, HealthCheckProtocol right) => left.Equals(right);
        public static bool operator !=(HealthCheckProtocol left, HealthCheckProtocol right) => !left.Equals(right);

        public static explicit operator string(HealthCheckProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HealthCheckProtocol other && Equals(other);
        public bool Equals(HealthCheckProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This field has been deprecated. You can no longer specify an algorithm for load balancers.
    /// </summary>
    [EnumType]
    public readonly struct LoadBalancerBaseAlgorithm : IEquatable<LoadBalancerBaseAlgorithm>
    {
        private readonly string _value;

        private LoadBalancerBaseAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LoadBalancerBaseAlgorithm RoundRobin { get; } = new LoadBalancerBaseAlgorithm("round_robin");
        public static LoadBalancerBaseAlgorithm LeastConnections { get; } = new LoadBalancerBaseAlgorithm("least_connections");

        public static bool operator ==(LoadBalancerBaseAlgorithm left, LoadBalancerBaseAlgorithm right) => left.Equals(right);
        public static bool operator !=(LoadBalancerBaseAlgorithm left, LoadBalancerBaseAlgorithm right) => !left.Equals(right);

        public static explicit operator string(LoadBalancerBaseAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LoadBalancerBaseAlgorithm other && Equals(other);
        public bool Equals(LoadBalancerBaseAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
    /// * `lb-small` = 1 node
    /// * `lb-medium` = 3 nodes
    /// * `lb-large` = 6 nodes
    /// 
    /// You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
    /// </summary>
    [EnumType]
    public readonly struct LoadBalancerBaseSize : IEquatable<LoadBalancerBaseSize>
    {
        private readonly string _value;

        private LoadBalancerBaseSize(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LoadBalancerBaseSize LbSmall { get; } = new LoadBalancerBaseSize("lb-small");
        public static LoadBalancerBaseSize LbMedium { get; } = new LoadBalancerBaseSize("lb-medium");
        public static LoadBalancerBaseSize LbLarge { get; } = new LoadBalancerBaseSize("lb-large");

        public static bool operator ==(LoadBalancerBaseSize left, LoadBalancerBaseSize right) => left.Equals(right);
        public static bool operator !=(LoadBalancerBaseSize left, LoadBalancerBaseSize right) => !left.Equals(right);

        public static explicit operator string(LoadBalancerBaseSize value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LoadBalancerBaseSize other && Equals(other);
        public bool Equals(LoadBalancerBaseSize other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
    /// </summary>
    [EnumType]
    public readonly struct LoadBalancerBaseStatus : IEquatable<LoadBalancerBaseStatus>
    {
        private readonly string _value;

        private LoadBalancerBaseStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LoadBalancerBaseStatus New { get; } = new LoadBalancerBaseStatus("new");
        public static LoadBalancerBaseStatus Active { get; } = new LoadBalancerBaseStatus("active");
        public static LoadBalancerBaseStatus Errored { get; } = new LoadBalancerBaseStatus("errored");

        public static bool operator ==(LoadBalancerBaseStatus left, LoadBalancerBaseStatus right) => left.Equals(right);
        public static bool operator !=(LoadBalancerBaseStatus left, LoadBalancerBaseStatus right) => !left.Equals(right);

        public static explicit operator string(LoadBalancerBaseStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LoadBalancerBaseStatus other && Equals(other);
        public bool Equals(LoadBalancerBaseStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`.
    /// </summary>
    [EnumType]
    public readonly struct StickySessionsType : IEquatable<StickySessionsType>
    {
        private readonly string _value;

        private StickySessionsType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StickySessionsType Cookies { get; } = new StickySessionsType("cookies");
        public static StickySessionsType None { get; } = new StickySessionsType("none");

        public static bool operator ==(StickySessionsType left, StickySessionsType right) => left.Equals(right);
        public static bool operator !=(StickySessionsType left, StickySessionsType right) => !left.Equals(right);

        public static explicit operator string(StickySessionsType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StickySessionsType other && Equals(other);
        public bool Equals(StickySessionsType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
