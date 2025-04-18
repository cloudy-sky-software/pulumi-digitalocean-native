# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from ._enums import *

__all__ = [
    'KubernetesNodePoolTaintArgs',
    'KubernetesNodePoolTaintArgsDict',
    'KubernetesNodePoolArgs',
    'KubernetesNodePoolArgsDict',
    'MaintenancePolicyArgs',
    'MaintenancePolicyArgsDict',
    'NodeStatusPropertiesArgs',
    'NodeStatusPropertiesArgsDict',
    'NodeArgs',
    'NodeArgsDict',
]

MYPY = False

if not MYPY:
    class KubernetesNodePoolTaintArgsDict(TypedDict):
        effect: NotRequired[pulumi.Input['KubernetesNodePoolTaintEffect']]
        """
        How the node reacts to pods that it won't tolerate. Available effect values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
        """
        key: NotRequired[pulumi.Input[builtins.str]]
        """
        An arbitrary string. The `key` and `value` fields of the `taint` object form a key-value pair. For example, if the value of the `key` field is "special" and the value of the `value` field is "gpu", the key value pair would be `special=gpu`.
        """
        value: NotRequired[pulumi.Input[builtins.str]]
        """
        An arbitrary string. The `key` and `value` fields of the `taint` object form a key-value pair. For example, if the value of the `key` field is "special" and the value of the `value` field is "gpu", the key value pair would be `special=gpu`.
        """
elif False:
    KubernetesNodePoolTaintArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class KubernetesNodePoolTaintArgs:
    def __init__(__self__, *,
                 effect: Optional[pulumi.Input['KubernetesNodePoolTaintEffect']] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input['KubernetesNodePoolTaintEffect'] effect: How the node reacts to pods that it won't tolerate. Available effect values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
        :param pulumi.Input[builtins.str] key: An arbitrary string. The `key` and `value` fields of the `taint` object form a key-value pair. For example, if the value of the `key` field is "special" and the value of the `value` field is "gpu", the key value pair would be `special=gpu`.
        :param pulumi.Input[builtins.str] value: An arbitrary string. The `key` and `value` fields of the `taint` object form a key-value pair. For example, if the value of the `key` field is "special" and the value of the `value` field is "gpu", the key value pair would be `special=gpu`.
        """
        if effect is not None:
            pulumi.set(__self__, "effect", effect)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def effect(self) -> Optional[pulumi.Input['KubernetesNodePoolTaintEffect']]:
        """
        How the node reacts to pods that it won't tolerate. Available effect values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
        """
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: Optional[pulumi.Input['KubernetesNodePoolTaintEffect']]):
        pulumi.set(self, "effect", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An arbitrary string. The `key` and `value` fields of the `taint` object form a key-value pair. For example, if the value of the `key` field is "special" and the value of the `value` field is "gpu", the key value pair would be `special=gpu`.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An arbitrary string. The `key` and `value` fields of the `taint` object form a key-value pair. For example, if the value of the `key` field is "special" and the value of the `value` field is "gpu", the key value pair would be `special=gpu`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class KubernetesNodePoolArgsDict(TypedDict):
        auto_scale: NotRequired[pulumi.Input[builtins.bool]]
        """
        A boolean value indicating whether auto-scaling is enabled for this node pool.
        """
        count: NotRequired[pulumi.Input[builtins.int]]
        """
        The number of Droplet instances in the node pool.
        """
        id: NotRequired[pulumi.Input[builtins.str]]
        """
        A unique ID that can be used to identify and reference a specific node pool.
        """
        labels: NotRequired[Any]
        """
        An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
        """
        max_nodes: NotRequired[pulumi.Input[builtins.int]]
        """
        The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
        """
        min_nodes: NotRequired[pulumi.Input[builtins.int]]
        """
        The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        A human-readable name for the node pool.
        """
        nodes: NotRequired[pulumi.Input[Sequence[pulumi.Input['NodeArgsDict']]]]
        """
        An object specifying the details of a specific worker node in a node pool.
        """
        size: NotRequired[pulumi.Input[builtins.str]]
        """
        The slug identifier for the type of Droplet used as workers in the node pool.
        """
        tags: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
        """
        taints: NotRequired[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolTaintArgsDict']]]]
        """
        An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
        """
elif False:
    KubernetesNodePoolArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class KubernetesNodePoolArgs:
    def __init__(__self__, *,
                 auto_scale: Optional[pulumi.Input[builtins.bool]] = None,
                 count: Optional[pulumi.Input[builtins.int]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[Any] = None,
                 max_nodes: Optional[pulumi.Input[builtins.int]] = None,
                 min_nodes: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['NodeArgs']]]] = None,
                 size: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 taints: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolTaintArgs']]]] = None):
        """
        :param pulumi.Input[builtins.bool] auto_scale: A boolean value indicating whether auto-scaling is enabled for this node pool.
        :param pulumi.Input[builtins.int] count: The number of Droplet instances in the node pool.
        :param pulumi.Input[builtins.str] id: A unique ID that can be used to identify and reference a specific node pool.
        :param Any labels: An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
        :param pulumi.Input[builtins.int] max_nodes: The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
        :param pulumi.Input[builtins.int] min_nodes: The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
        :param pulumi.Input[builtins.str] name: A human-readable name for the node pool.
        :param pulumi.Input[Sequence[pulumi.Input['NodeArgs']]] nodes: An object specifying the details of a specific worker node in a node pool.
        :param pulumi.Input[builtins.str] size: The slug identifier for the type of Droplet used as workers in the node pool.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolTaintArgs']]] taints: An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
        """
        if auto_scale is not None:
            pulumi.set(__self__, "auto_scale", auto_scale)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if max_nodes is not None:
            pulumi.set(__self__, "max_nodes", max_nodes)
        if min_nodes is not None:
            pulumi.set(__self__, "min_nodes", min_nodes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if taints is not None:
            pulumi.set(__self__, "taints", taints)

    @property
    @pulumi.getter(name="autoScale")
    def auto_scale(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A boolean value indicating whether auto-scaling is enabled for this node pool.
        """
        return pulumi.get(self, "auto_scale")

    @auto_scale.setter
    def auto_scale(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "auto_scale", value)

    @property
    @pulumi.getter
    def count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of Droplet instances in the node pool.
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique ID that can be used to identify and reference a specific node pool.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[Any]:
        """
        An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[Any]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="maxNodes")
    def max_nodes(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
        """
        return pulumi.get(self, "max_nodes")

    @max_nodes.setter
    def max_nodes(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_nodes", value)

    @property
    @pulumi.getter(name="minNodes")
    def min_nodes(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
        """
        return pulumi.get(self, "min_nodes")

    @min_nodes.setter
    def min_nodes(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_nodes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A human-readable name for the node pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NodeArgs']]]]:
        """
        An object specifying the details of a specific worker node in a node pool.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The slug identifier for the type of Droplet used as workers in the node pool.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def taints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolTaintArgs']]]]:
        """
        An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
        """
        return pulumi.get(self, "taints")

    @taints.setter
    def taints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolTaintArgs']]]]):
        pulumi.set(self, "taints", value)


if not MYPY:
    class MaintenancePolicyArgsDict(TypedDict):
        """
        An object specifying the maintenance window policy for the Kubernetes cluster.
        """
        day: NotRequired[pulumi.Input['MaintenancePolicyDay']]
        """
        The day of the maintenance window policy. May be one of `monday` through `sunday`, or `any` to indicate an arbitrary week day.
        """
        duration: NotRequired[pulumi.Input[builtins.str]]
        """
        The duration of the maintenance window policy in human-readable format.
        """
        start_time: NotRequired[pulumi.Input[builtins.str]]
        """
        The start time in UTC of the maintenance window policy in 24-hour clock format / HH:MM notation (e.g., `15:00`).
        """
elif False:
    MaintenancePolicyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MaintenancePolicyArgs:
    def __init__(__self__, *,
                 day: Optional[pulumi.Input['MaintenancePolicyDay']] = None,
                 duration: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.str]] = None):
        """
        An object specifying the maintenance window policy for the Kubernetes cluster.
        :param pulumi.Input['MaintenancePolicyDay'] day: The day of the maintenance window policy. May be one of `monday` through `sunday`, or `any` to indicate an arbitrary week day.
        :param pulumi.Input[builtins.str] duration: The duration of the maintenance window policy in human-readable format.
        :param pulumi.Input[builtins.str] start_time: The start time in UTC of the maintenance window policy in 24-hour clock format / HH:MM notation (e.g., `15:00`).
        """
        if day is not None:
            pulumi.set(__self__, "day", day)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter
    def day(self) -> Optional[pulumi.Input['MaintenancePolicyDay']]:
        """
        The day of the maintenance window policy. May be one of `monday` through `sunday`, or `any` to indicate an arbitrary week day.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: Optional[pulumi.Input['MaintenancePolicyDay']]):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The duration of the maintenance window policy in human-readable format.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The start time in UTC of the maintenance window policy in 24-hour clock format / HH:MM notation (e.g., `15:00`).
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "start_time", value)


if not MYPY:
    class NodeStatusPropertiesArgsDict(TypedDict):
        """
        An object containing a `state` attribute whose value is set to a string indicating the current status of the node.
        """
        state: NotRequired[pulumi.Input['NodeStatusPropertiesState']]
        """
        A string indicating the current status of the node.
        """
elif False:
    NodeStatusPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class NodeStatusPropertiesArgs:
    def __init__(__self__, *,
                 state: Optional[pulumi.Input['NodeStatusPropertiesState']] = None):
        """
        An object containing a `state` attribute whose value is set to a string indicating the current status of the node.
        :param pulumi.Input['NodeStatusPropertiesState'] state: A string indicating the current status of the node.
        """
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['NodeStatusPropertiesState']]:
        """
        A string indicating the current status of the node.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['NodeStatusPropertiesState']]):
        pulumi.set(self, "state", value)


if not MYPY:
    class NodeArgsDict(TypedDict):
        created_at: NotRequired[pulumi.Input[builtins.str]]
        """
        A time value given in ISO8601 combined date and time format that represents when the node was created.
        """
        droplet_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The ID of the Droplet used for the worker node.
        """
        id: NotRequired[pulumi.Input[builtins.str]]
        """
        A unique ID that can be used to identify and reference the node.
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        An automatically generated, human-readable name for the node.
        """
        status: NotRequired[pulumi.Input['NodeStatusPropertiesArgsDict']]
        """
        An object containing a `state` attribute whose value is set to a string indicating the current status of the node.
        """
        updated_at: NotRequired[pulumi.Input[builtins.str]]
        """
        A time value given in ISO8601 combined date and time format that represents when the node was last updated.
        """
elif False:
    NodeArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class NodeArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 droplet_id: Optional[pulumi.Input[builtins.str]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input['NodeStatusPropertiesArgs']] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] created_at: A time value given in ISO8601 combined date and time format that represents when the node was created.
        :param pulumi.Input[builtins.str] droplet_id: The ID of the Droplet used for the worker node.
        :param pulumi.Input[builtins.str] id: A unique ID that can be used to identify and reference the node.
        :param pulumi.Input[builtins.str] name: An automatically generated, human-readable name for the node.
        :param pulumi.Input['NodeStatusPropertiesArgs'] status: An object containing a `state` attribute whose value is set to a string indicating the current status of the node.
        :param pulumi.Input[builtins.str] updated_at: A time value given in ISO8601 combined date and time format that represents when the node was last updated.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if droplet_id is not None:
            pulumi.set(__self__, "droplet_id", droplet_id)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the node was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the Droplet used for the worker node.
        """
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "droplet_id", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique ID that can be used to identify and reference the node.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An automatically generated, human-readable name for the node.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['NodeStatusPropertiesArgs']]:
        """
        An object containing a `state` attribute whose value is set to a string indicating the current status of the node.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['NodeStatusPropertiesArgs']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the node was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


