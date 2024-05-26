// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.DigitalOceanNative.VolumesV2
{
    /// <summary>
    /// The current status of the action. This can be "in-progress", "completed", or "errored".
    /// </summary>
    [EnumType]
    public readonly struct ActionStatus : IEquatable<ActionStatus>
    {
        private readonly string _value;

        private ActionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ActionStatus InProgress { get; } = new ActionStatus("in-progress");
        public static ActionStatus Completed { get; } = new ActionStatus("completed");
        public static ActionStatus Errored { get; } = new ActionStatus("errored");

        public static bool operator ==(ActionStatus left, ActionStatus right) => left.Equals(right);
        public static bool operator !=(ActionStatus left, ActionStatus right) => !left.Equals(right);

        public static explicit operator string(ActionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ActionStatus other && Equals(other);
        public bool Equals(ActionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The slug identifier for the region where the resource will initially be  available.
    /// </summary>
    [EnumType]
    public readonly struct Ext4PropertiesRegion : IEquatable<Ext4PropertiesRegion>
    {
        private readonly string _value;

        private Ext4PropertiesRegion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Ext4PropertiesRegion Ams1 { get; } = new Ext4PropertiesRegion("ams1");
        public static Ext4PropertiesRegion Ams2 { get; } = new Ext4PropertiesRegion("ams2");
        public static Ext4PropertiesRegion Ams3 { get; } = new Ext4PropertiesRegion("ams3");
        public static Ext4PropertiesRegion Blr1 { get; } = new Ext4PropertiesRegion("blr1");
        public static Ext4PropertiesRegion Fra1 { get; } = new Ext4PropertiesRegion("fra1");
        public static Ext4PropertiesRegion Lon1 { get; } = new Ext4PropertiesRegion("lon1");
        public static Ext4PropertiesRegion Nyc1 { get; } = new Ext4PropertiesRegion("nyc1");
        public static Ext4PropertiesRegion Nyc2 { get; } = new Ext4PropertiesRegion("nyc2");
        public static Ext4PropertiesRegion Nyc3 { get; } = new Ext4PropertiesRegion("nyc3");
        public static Ext4PropertiesRegion Sfo1 { get; } = new Ext4PropertiesRegion("sfo1");
        public static Ext4PropertiesRegion Sfo2 { get; } = new Ext4PropertiesRegion("sfo2");
        public static Ext4PropertiesRegion Sfo3 { get; } = new Ext4PropertiesRegion("sfo3");
        public static Ext4PropertiesRegion Sgp1 { get; } = new Ext4PropertiesRegion("sgp1");
        public static Ext4PropertiesRegion Tor1 { get; } = new Ext4PropertiesRegion("tor1");

        public static bool operator ==(Ext4PropertiesRegion left, Ext4PropertiesRegion right) => left.Equals(right);
        public static bool operator !=(Ext4PropertiesRegion left, Ext4PropertiesRegion right) => !left.Equals(right);

        public static explicit operator string(Ext4PropertiesRegion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Ext4PropertiesRegion other && Equals(other);
        public bool Equals(Ext4PropertiesRegion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of resource that the snapshot originated from.
    /// </summary>
    [EnumType]
    public readonly struct SnapshotsPropertiesResourceType : IEquatable<SnapshotsPropertiesResourceType>
    {
        private readonly string _value;

        private SnapshotsPropertiesResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SnapshotsPropertiesResourceType Droplet { get; } = new SnapshotsPropertiesResourceType("droplet");
        public static SnapshotsPropertiesResourceType Volume { get; } = new SnapshotsPropertiesResourceType("volume");

        public static bool operator ==(SnapshotsPropertiesResourceType left, SnapshotsPropertiesResourceType right) => left.Equals(right);
        public static bool operator !=(SnapshotsPropertiesResourceType left, SnapshotsPropertiesResourceType right) => !left.Equals(right);

        public static explicit operator string(SnapshotsPropertiesResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SnapshotsPropertiesResourceType other && Equals(other);
        public bool Equals(SnapshotsPropertiesResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The slug identifier for the region where the resource will initially be  available.
    /// </summary>
    [EnumType]
    public readonly struct VolumeActionCreateBaseRegion : IEquatable<VolumeActionCreateBaseRegion>
    {
        private readonly string _value;

        private VolumeActionCreateBaseRegion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VolumeActionCreateBaseRegion Ams1 { get; } = new VolumeActionCreateBaseRegion("ams1");
        public static VolumeActionCreateBaseRegion Ams2 { get; } = new VolumeActionCreateBaseRegion("ams2");
        public static VolumeActionCreateBaseRegion Ams3 { get; } = new VolumeActionCreateBaseRegion("ams3");
        public static VolumeActionCreateBaseRegion Blr1 { get; } = new VolumeActionCreateBaseRegion("blr1");
        public static VolumeActionCreateBaseRegion Fra1 { get; } = new VolumeActionCreateBaseRegion("fra1");
        public static VolumeActionCreateBaseRegion Lon1 { get; } = new VolumeActionCreateBaseRegion("lon1");
        public static VolumeActionCreateBaseRegion Nyc1 { get; } = new VolumeActionCreateBaseRegion("nyc1");
        public static VolumeActionCreateBaseRegion Nyc2 { get; } = new VolumeActionCreateBaseRegion("nyc2");
        public static VolumeActionCreateBaseRegion Nyc3 { get; } = new VolumeActionCreateBaseRegion("nyc3");
        public static VolumeActionCreateBaseRegion Sfo1 { get; } = new VolumeActionCreateBaseRegion("sfo1");
        public static VolumeActionCreateBaseRegion Sfo2 { get; } = new VolumeActionCreateBaseRegion("sfo2");
        public static VolumeActionCreateBaseRegion Sfo3 { get; } = new VolumeActionCreateBaseRegion("sfo3");
        public static VolumeActionCreateBaseRegion Sgp1 { get; } = new VolumeActionCreateBaseRegion("sgp1");
        public static VolumeActionCreateBaseRegion Tor1 { get; } = new VolumeActionCreateBaseRegion("tor1");

        public static bool operator ==(VolumeActionCreateBaseRegion left, VolumeActionCreateBaseRegion right) => left.Equals(right);
        public static bool operator !=(VolumeActionCreateBaseRegion left, VolumeActionCreateBaseRegion right) => !left.Equals(right);

        public static explicit operator string(VolumeActionCreateBaseRegion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeActionCreateBaseRegion other && Equals(other);
        public bool Equals(VolumeActionCreateBaseRegion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The volume action to initiate.
    /// </summary>
    [EnumType]
    public readonly struct VolumeActionCreateBaseType : IEquatable<VolumeActionCreateBaseType>
    {
        private readonly string _value;

        private VolumeActionCreateBaseType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VolumeActionCreateBaseType Attach { get; } = new VolumeActionCreateBaseType("attach");
        public static VolumeActionCreateBaseType Detach { get; } = new VolumeActionCreateBaseType("detach");
        public static VolumeActionCreateBaseType Resize { get; } = new VolumeActionCreateBaseType("resize");

        public static bool operator ==(VolumeActionCreateBaseType left, VolumeActionCreateBaseType right) => left.Equals(right);
        public static bool operator !=(VolumeActionCreateBaseType left, VolumeActionCreateBaseType right) => !left.Equals(right);

        public static explicit operator string(VolumeActionCreateBaseType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeActionCreateBaseType other && Equals(other);
        public bool Equals(VolumeActionCreateBaseType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The slug identifier for the region where the resource will initially be  available.
    /// </summary>
    [EnumType]
    public readonly struct XfsPropertiesRegion : IEquatable<XfsPropertiesRegion>
    {
        private readonly string _value;

        private XfsPropertiesRegion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static XfsPropertiesRegion Ams1 { get; } = new XfsPropertiesRegion("ams1");
        public static XfsPropertiesRegion Ams2 { get; } = new XfsPropertiesRegion("ams2");
        public static XfsPropertiesRegion Ams3 { get; } = new XfsPropertiesRegion("ams3");
        public static XfsPropertiesRegion Blr1 { get; } = new XfsPropertiesRegion("blr1");
        public static XfsPropertiesRegion Fra1 { get; } = new XfsPropertiesRegion("fra1");
        public static XfsPropertiesRegion Lon1 { get; } = new XfsPropertiesRegion("lon1");
        public static XfsPropertiesRegion Nyc1 { get; } = new XfsPropertiesRegion("nyc1");
        public static XfsPropertiesRegion Nyc2 { get; } = new XfsPropertiesRegion("nyc2");
        public static XfsPropertiesRegion Nyc3 { get; } = new XfsPropertiesRegion("nyc3");
        public static XfsPropertiesRegion Sfo1 { get; } = new XfsPropertiesRegion("sfo1");
        public static XfsPropertiesRegion Sfo2 { get; } = new XfsPropertiesRegion("sfo2");
        public static XfsPropertiesRegion Sfo3 { get; } = new XfsPropertiesRegion("sfo3");
        public static XfsPropertiesRegion Sgp1 { get; } = new XfsPropertiesRegion("sgp1");
        public static XfsPropertiesRegion Tor1 { get; } = new XfsPropertiesRegion("tor1");

        public static bool operator ==(XfsPropertiesRegion left, XfsPropertiesRegion right) => left.Equals(right);
        public static bool operator !=(XfsPropertiesRegion left, XfsPropertiesRegion right) => !left.Equals(right);

        public static explicit operator string(XfsPropertiesRegion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is XfsPropertiesRegion other && Equals(other);
        public bool Equals(XfsPropertiesRegion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
