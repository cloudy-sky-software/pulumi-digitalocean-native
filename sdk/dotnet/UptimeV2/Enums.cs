// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.DigitalOceanNative.UptimeV2
{
    /// <summary>
    /// The comparison operator used against the alert's threshold.
    /// </summary>
    [EnumType]
    public readonly struct AlertUpdatableComparison : IEquatable<AlertUpdatableComparison>
    {
        private readonly string _value;

        private AlertUpdatableComparison(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertUpdatableComparison GreaterThan { get; } = new AlertUpdatableComparison("greater_than");
        public static AlertUpdatableComparison LessThan { get; } = new AlertUpdatableComparison("less_than");

        public static bool operator ==(AlertUpdatableComparison left, AlertUpdatableComparison right) => left.Equals(right);
        public static bool operator !=(AlertUpdatableComparison left, AlertUpdatableComparison right) => !left.Equals(right);

        public static explicit operator string(AlertUpdatableComparison value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertUpdatableComparison other && Equals(other);
        public bool Equals(AlertUpdatableComparison other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Period of time the threshold must be exceeded to trigger the alert.
    /// </summary>
    [EnumType]
    public readonly struct AlertUpdatablePeriod : IEquatable<AlertUpdatablePeriod>
    {
        private readonly string _value;

        private AlertUpdatablePeriod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertUpdatablePeriod AlertUpdatablePeriod_2m { get; } = new AlertUpdatablePeriod("2m");
        public static AlertUpdatablePeriod AlertUpdatablePeriod_3m { get; } = new AlertUpdatablePeriod("3m");
        public static AlertUpdatablePeriod AlertUpdatablePeriod_5m { get; } = new AlertUpdatablePeriod("5m");
        public static AlertUpdatablePeriod AlertUpdatablePeriod_10m { get; } = new AlertUpdatablePeriod("10m");
        public static AlertUpdatablePeriod AlertUpdatablePeriod_15m { get; } = new AlertUpdatablePeriod("15m");
        public static AlertUpdatablePeriod AlertUpdatablePeriod_30m { get; } = new AlertUpdatablePeriod("30m");
        public static AlertUpdatablePeriod AlertUpdatablePeriod_1h { get; } = new AlertUpdatablePeriod("1h");

        public static bool operator ==(AlertUpdatablePeriod left, AlertUpdatablePeriod right) => left.Equals(right);
        public static bool operator !=(AlertUpdatablePeriod left, AlertUpdatablePeriod right) => !left.Equals(right);

        public static explicit operator string(AlertUpdatablePeriod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertUpdatablePeriod other && Equals(other);
        public bool Equals(AlertUpdatablePeriod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of alert.
    /// </summary>
    [EnumType]
    public readonly struct AlertUpdatableType : IEquatable<AlertUpdatableType>
    {
        private readonly string _value;

        private AlertUpdatableType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertUpdatableType Latency { get; } = new AlertUpdatableType("latency");
        public static AlertUpdatableType Down { get; } = new AlertUpdatableType("down");
        public static AlertUpdatableType DownGlobal { get; } = new AlertUpdatableType("down_global");
        public static AlertUpdatableType SslExpiry { get; } = new AlertUpdatableType("ssl_expiry");

        public static bool operator ==(AlertUpdatableType left, AlertUpdatableType right) => left.Equals(right);
        public static bool operator !=(AlertUpdatableType left, AlertUpdatableType right) => !left.Equals(right);

        public static explicit operator string(AlertUpdatableType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertUpdatableType other && Equals(other);
        public bool Equals(AlertUpdatableType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CheckUpdatableRegionsItem : IEquatable<CheckUpdatableRegionsItem>
    {
        private readonly string _value;

        private CheckUpdatableRegionsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CheckUpdatableRegionsItem UsEast { get; } = new CheckUpdatableRegionsItem("us_east");
        public static CheckUpdatableRegionsItem UsWest { get; } = new CheckUpdatableRegionsItem("us_west");
        public static CheckUpdatableRegionsItem EuWest { get; } = new CheckUpdatableRegionsItem("eu_west");
        public static CheckUpdatableRegionsItem SeAsia { get; } = new CheckUpdatableRegionsItem("se_asia");

        public static bool operator ==(CheckUpdatableRegionsItem left, CheckUpdatableRegionsItem right) => left.Equals(right);
        public static bool operator !=(CheckUpdatableRegionsItem left, CheckUpdatableRegionsItem right) => !left.Equals(right);

        public static explicit operator string(CheckUpdatableRegionsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CheckUpdatableRegionsItem other && Equals(other);
        public bool Equals(CheckUpdatableRegionsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of health check to perform.
    /// </summary>
    [EnumType]
    public readonly struct CheckUpdatableType : IEquatable<CheckUpdatableType>
    {
        private readonly string _value;

        private CheckUpdatableType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CheckUpdatableType Ping { get; } = new CheckUpdatableType("ping");
        public static CheckUpdatableType Http { get; } = new CheckUpdatableType("http");
        public static CheckUpdatableType Https { get; } = new CheckUpdatableType("https");

        public static bool operator ==(CheckUpdatableType left, CheckUpdatableType right) => left.Equals(right);
        public static bool operator !=(CheckUpdatableType left, CheckUpdatableType right) => !left.Equals(right);

        public static explicit operator string(CheckUpdatableType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CheckUpdatableType other && Equals(other);
        public bool Equals(CheckUpdatableType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct RegionStateStatus : IEquatable<RegionStateStatus>
    {
        private readonly string _value;

        private RegionStateStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RegionStateStatus Down { get; } = new RegionStateStatus("DOWN");
        public static RegionStateStatus Up { get; } = new RegionStateStatus("UP");
        public static RegionStateStatus Checking { get; } = new RegionStateStatus("CHECKING");

        public static bool operator ==(RegionStateStatus left, RegionStateStatus right) => left.Equals(right);
        public static bool operator !=(RegionStateStatus left, RegionStateStatus right) => !left.Equals(right);

        public static explicit operator string(RegionStateStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RegionStateStatus other && Equals(other);
        public bool Equals(RegionStateStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
