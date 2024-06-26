// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ActionsV2
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
}
