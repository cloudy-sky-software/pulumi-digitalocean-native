# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['FirewallsArgs', 'Firewalls']

@pulumi.input_type
class FirewallsArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 droplet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesInboundRulesItemArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesOutboundRulesItemArgs']]]] = None,
                 pending_changes: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallPropertiesPendingChangesItemPropertiesArgs']]]] = None,
                 status: Optional[pulumi.Input['FirewallPropertiesStatus']] = None,
                 tags: Optional[pulumi.Input['FirewallPropertiesTagsArgs']] = None):
        """
        The set of arguments for constructing a Firewalls resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the firewall was created.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] droplet_ids: An array containing the IDs of the Droplets assigned to the firewall.
        :param pulumi.Input[str] name: A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
        :param pulumi.Input[Sequence[pulumi.Input['FirewallPropertiesPendingChangesItemPropertiesArgs']]] pending_changes: An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
        :param pulumi.Input['FirewallPropertiesStatus'] status: A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if droplet_ids is not None:
            pulumi.set(__self__, "droplet_ids", droplet_ids)
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)
        if pending_changes is not None:
            pulumi.set(__self__, "pending_changes", pending_changes)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the firewall was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dropletIds")
    def droplet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        An array containing the IDs of the Droplets assigned to the firewall.
        """
        return pulumi.get(self, "droplet_ids")

    @droplet_ids.setter
    def droplet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "droplet_ids", value)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesInboundRulesItemArgs']]]]:
        return pulumi.get(self, "inbound_rules")

    @inbound_rules.setter
    def inbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesInboundRulesItemArgs']]]]):
        pulumi.set(self, "inbound_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesOutboundRulesItemArgs']]]]:
        return pulumi.get(self, "outbound_rules")

    @outbound_rules.setter
    def outbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesOutboundRulesItemArgs']]]]):
        pulumi.set(self, "outbound_rules", value)

    @property
    @pulumi.getter(name="pendingChanges")
    def pending_changes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallPropertiesPendingChangesItemPropertiesArgs']]]]:
        """
        An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
        """
        return pulumi.get(self, "pending_changes")

    @pending_changes.setter
    def pending_changes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallPropertiesPendingChangesItemPropertiesArgs']]]]):
        pulumi.set(self, "pending_changes", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['FirewallPropertiesStatus']]:
        """
        A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['FirewallPropertiesStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['FirewallPropertiesTagsArgs']]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['FirewallPropertiesTagsArgs']]):
        pulumi.set(self, "tags", value)


class Firewalls(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 droplet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallRulesInboundRulesItemArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallRulesOutboundRulesItemArgs']]]]] = None,
                 pending_changes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallPropertiesPendingChangesItemPropertiesArgs']]]]] = None,
                 status: Optional[pulumi.Input['FirewallPropertiesStatus']] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['FirewallPropertiesTagsArgs']]] = None,
                 __props__=None):
        """
        Create a Firewalls resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the firewall was created.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] droplet_ids: An array containing the IDs of the Droplets assigned to the firewall.
        :param pulumi.Input[str] name: A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallPropertiesPendingChangesItemPropertiesArgs']]]] pending_changes: An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
        :param pulumi.Input['FirewallPropertiesStatus'] status: A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Firewalls resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 droplet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallRulesInboundRulesItemArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallRulesOutboundRulesItemArgs']]]]] = None,
                 pending_changes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallPropertiesPendingChangesItemPropertiesArgs']]]]] = None,
                 status: Optional[pulumi.Input['FirewallPropertiesStatus']] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['FirewallPropertiesTagsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallsArgs.__new__(FirewallsArgs)

            __props__.__dict__["created_at"] = created_at
            __props__.__dict__["droplet_ids"] = droplet_ids
            __props__.__dict__["inbound_rules"] = inbound_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["outbound_rules"] = outbound_rules
            __props__.__dict__["pending_changes"] = pending_changes
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["firewall"] = None
        super(Firewalls, __self__).__init__(
            'digitalocean-native:firewalls/v2:Firewalls',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Firewalls':
        """
        Get an existing Firewalls resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FirewallsArgs.__new__(FirewallsArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["droplet_ids"] = None
        __props__.__dict__["firewall"] = None
        __props__.__dict__["inbound_rules"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["outbound_rules"] = None
        __props__.__dict__["pending_changes"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        return Firewalls(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the firewall was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dropletIds")
    def droplet_ids(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        An array containing the IDs of the Droplets assigned to the firewall.
        """
        return pulumi.get(self, "droplet_ids")

    @property
    @pulumi.getter
    def firewall(self) -> pulumi.Output[Optional['outputs.Firewall']]:
        return pulumi.get(self, "firewall")

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallRulesInboundRulesItem']]]:
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallRulesOutboundRulesItem']]]:
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="pendingChanges")
    def pending_changes(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallPropertiesPendingChangesItemProperties']]]:
        """
        An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
        """
        return pulumi.get(self, "pending_changes")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['FirewallPropertiesStatus']]:
        """
        A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional['outputs.FirewallPropertiesTags']]:
        return pulumi.get(self, "tags")
