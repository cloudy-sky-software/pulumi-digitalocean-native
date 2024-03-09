// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.DigitalOceanNative.KubernetesV2
{
    /// <summary>
    /// A string indicating the current status of the cluster.
    /// </summary>
    [EnumType]
    public readonly struct ClusterStatusPropertiesState : IEquatable<ClusterStatusPropertiesState>
    {
        private readonly string _value;

        private ClusterStatusPropertiesState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterStatusPropertiesState Running { get; } = new ClusterStatusPropertiesState("running");
        public static ClusterStatusPropertiesState Provisioning { get; } = new ClusterStatusPropertiesState("provisioning");
        public static ClusterStatusPropertiesState Degraded { get; } = new ClusterStatusPropertiesState("degraded");
        public static ClusterStatusPropertiesState Error { get; } = new ClusterStatusPropertiesState("error");
        public static ClusterStatusPropertiesState Deleted { get; } = new ClusterStatusPropertiesState("deleted");
        public static ClusterStatusPropertiesState Upgrading { get; } = new ClusterStatusPropertiesState("upgrading");
        public static ClusterStatusPropertiesState Deleting { get; } = new ClusterStatusPropertiesState("deleting");

        public static bool operator ==(ClusterStatusPropertiesState left, ClusterStatusPropertiesState right) => left.Equals(right);
        public static bool operator !=(ClusterStatusPropertiesState left, ClusterStatusPropertiesState right) => !left.Equals(right);

        public static explicit operator string(ClusterStatusPropertiesState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterStatusPropertiesState other && Equals(other);
        public bool Equals(ClusterStatusPropertiesState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// How the node reacts to pods that it won't tolerate. Available effect values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
    /// </summary>
    [EnumType]
    public readonly struct KubernetesNodePoolTaintEffect : IEquatable<KubernetesNodePoolTaintEffect>
    {
        private readonly string _value;

        private KubernetesNodePoolTaintEffect(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KubernetesNodePoolTaintEffect NoSchedule { get; } = new KubernetesNodePoolTaintEffect("NoSchedule");
        public static KubernetesNodePoolTaintEffect PreferNoSchedule { get; } = new KubernetesNodePoolTaintEffect("PreferNoSchedule");
        public static KubernetesNodePoolTaintEffect NoExecute { get; } = new KubernetesNodePoolTaintEffect("NoExecute");

        public static bool operator ==(KubernetesNodePoolTaintEffect left, KubernetesNodePoolTaintEffect right) => left.Equals(right);
        public static bool operator !=(KubernetesNodePoolTaintEffect left, KubernetesNodePoolTaintEffect right) => !left.Equals(right);

        public static explicit operator string(KubernetesNodePoolTaintEffect value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KubernetesNodePoolTaintEffect other && Equals(other);
        public bool Equals(KubernetesNodePoolTaintEffect other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The day of the maintenance window policy. May be one of `monday` through `sunday`, or `any` to indicate an arbitrary week day.
    /// </summary>
    [EnumType]
    public readonly struct MaintenancePolicyDay : IEquatable<MaintenancePolicyDay>
    {
        private readonly string _value;

        private MaintenancePolicyDay(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MaintenancePolicyDay Any { get; } = new MaintenancePolicyDay("any");
        public static MaintenancePolicyDay Monday { get; } = new MaintenancePolicyDay("monday");
        public static MaintenancePolicyDay Tuesday { get; } = new MaintenancePolicyDay("tuesday");
        public static MaintenancePolicyDay Wednesday { get; } = new MaintenancePolicyDay("wednesday");
        public static MaintenancePolicyDay Thursday { get; } = new MaintenancePolicyDay("thursday");
        public static MaintenancePolicyDay Friday { get; } = new MaintenancePolicyDay("friday");
        public static MaintenancePolicyDay Saturday { get; } = new MaintenancePolicyDay("saturday");
        public static MaintenancePolicyDay Sunday { get; } = new MaintenancePolicyDay("sunday");

        public static bool operator ==(MaintenancePolicyDay left, MaintenancePolicyDay right) => left.Equals(right);
        public static bool operator !=(MaintenancePolicyDay left, MaintenancePolicyDay right) => !left.Equals(right);

        public static explicit operator string(MaintenancePolicyDay value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MaintenancePolicyDay other && Equals(other);
        public bool Equals(MaintenancePolicyDay other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A string indicating the current status of the node.
    /// </summary>
    [EnumType]
    public readonly struct NodeStatusPropertiesState : IEquatable<NodeStatusPropertiesState>
    {
        private readonly string _value;

        private NodeStatusPropertiesState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NodeStatusPropertiesState Provisioning { get; } = new NodeStatusPropertiesState("provisioning");
        public static NodeStatusPropertiesState Running { get; } = new NodeStatusPropertiesState("running");
        public static NodeStatusPropertiesState Draining { get; } = new NodeStatusPropertiesState("draining");
        public static NodeStatusPropertiesState Deleting { get; } = new NodeStatusPropertiesState("deleting");

        public static bool operator ==(NodeStatusPropertiesState left, NodeStatusPropertiesState right) => left.Equals(right);
        public static bool operator !=(NodeStatusPropertiesState left, NodeStatusPropertiesState right) => !left.Equals(right);

        public static explicit operator string(NodeStatusPropertiesState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NodeStatusPropertiesState other && Equals(other);
        public bool Equals(NodeStatusPropertiesState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A string indicating the current status of the cluster.
    /// </summary>
    [EnumType]
    public readonly struct StatusPropertiesState : IEquatable<StatusPropertiesState>
    {
        private readonly string _value;

        private StatusPropertiesState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StatusPropertiesState Running { get; } = new StatusPropertiesState("running");
        public static StatusPropertiesState Provisioning { get; } = new StatusPropertiesState("provisioning");
        public static StatusPropertiesState Degraded { get; } = new StatusPropertiesState("degraded");
        public static StatusPropertiesState Error { get; } = new StatusPropertiesState("error");
        public static StatusPropertiesState Deleted { get; } = new StatusPropertiesState("deleted");
        public static StatusPropertiesState Upgrading { get; } = new StatusPropertiesState("upgrading");
        public static StatusPropertiesState Deleting { get; } = new StatusPropertiesState("deleting");

        public static bool operator ==(StatusPropertiesState left, StatusPropertiesState right) => left.Equals(right);
        public static bool operator !=(StatusPropertiesState left, StatusPropertiesState right) => !left.Equals(right);

        public static explicit operator string(StatusPropertiesState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StatusPropertiesState other && Equals(other);
        public bool Equals(StatusPropertiesState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}