# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'ClusterStatusPropertiesState',
    'KubernetesNodePoolTaintEffect',
    'MaintenancePolicyDay',
    'NodeStatusPropertiesState',
    'StatusPropertiesState',
]


class ClusterStatusPropertiesState(builtins.str, Enum):
    """
    A string indicating the current status of the cluster.
    """
    RUNNING = "running"
    PROVISIONING = "provisioning"
    DEGRADED = "degraded"
    ERROR = "error"
    DELETED = "deleted"
    UPGRADING = "upgrading"
    DELETING = "deleting"


class KubernetesNodePoolTaintEffect(builtins.str, Enum):
    """
    How the node reacts to pods that it won't tolerate. Available effect values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
    """
    NO_SCHEDULE = "NoSchedule"
    PREFER_NO_SCHEDULE = "PreferNoSchedule"
    NO_EXECUTE = "NoExecute"


class MaintenancePolicyDay(builtins.str, Enum):
    """
    The day of the maintenance window policy. May be one of `monday` through `sunday`, or `any` to indicate an arbitrary week day.
    """
    ANY = "any"
    MONDAY = "monday"
    TUESDAY = "tuesday"
    WEDNESDAY = "wednesday"
    THURSDAY = "thursday"
    FRIDAY = "friday"
    SATURDAY = "saturday"
    SUNDAY = "sunday"


class NodeStatusPropertiesState(builtins.str, Enum):
    """
    A string indicating the current status of the node.
    """
    PROVISIONING = "provisioning"
    RUNNING = "running"
    DRAINING = "draining"
    DELETING = "deleting"


class StatusPropertiesState(builtins.str, Enum):
    """
    A string indicating the current status of the cluster.
    """
    RUNNING = "running"
    PROVISIONING = "provisioning"
    DEGRADED = "degraded"
    ERROR = "error"
    DELETED = "deleted"
    UPGRADING = "upgrading"
    DELETING = "deleting"
