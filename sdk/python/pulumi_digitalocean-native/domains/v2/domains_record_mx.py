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

__all__ = ['DomainsRecordMxArgs', 'DomainsRecordMx']

@pulumi.input_type
class DomainsRecordMxArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 data: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 flags: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a DomainsRecordMx resource.
        :param pulumi.Input[str] type: The type of the DNS record. For example: A, CNAME, TXT, ...
        :param pulumi.Input[str] data: Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
        :param pulumi.Input[str] domain_name: The name of the domain itself.
        :param pulumi.Input[int] flags: An unsigned integer between 0-255 used for CAA records.
        :param pulumi.Input[str] name: The host name, alias, or service being defined by the record.
        :param pulumi.Input[int] port: The port for SRV records.
        :param pulumi.Input[int] priority: The priority for SRV and MX records.
        :param pulumi.Input[str] tag: The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
        :param pulumi.Input[int] ttl: This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
        :param pulumi.Input[int] weight: The weight for SRV records.
        """
        pulumi.set(__self__, "type", type)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if flags is not None:
            pulumi.set(__self__, "flags", flags)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the DNS record. For example: A, CNAME, TXT, ...
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the domain itself.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def flags(self) -> Optional[pulumi.Input[int]]:
        """
        An unsigned integer between 0-255 used for CAA records.
        """
        return pulumi.get(self, "flags")

    @flags.setter
    def flags(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "flags", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host name, alias, or service being defined by the record.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port for SRV records.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority for SRV and MX records.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight for SRV records.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class DomainsRecordMx(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 flags: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a DomainsRecordMx resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
        :param pulumi.Input[str] domain_name: The name of the domain itself.
        :param pulumi.Input[int] flags: An unsigned integer between 0-255 used for CAA records.
        :param pulumi.Input[str] name: The host name, alias, or service being defined by the record.
        :param pulumi.Input[int] port: The port for SRV records.
        :param pulumi.Input[int] priority: The priority for SRV and MX records.
        :param pulumi.Input[str] tag: The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
        :param pulumi.Input[int] ttl: This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
        :param pulumi.Input[str] type: The type of the DNS record. For example: A, CNAME, TXT, ...
        :param pulumi.Input[int] weight: The weight for SRV records.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainsRecordMxArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DomainsRecordMx resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DomainsRecordMxArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainsRecordMxArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 flags: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainsRecordMxArgs.__new__(DomainsRecordMxArgs)

            __props__.__dict__["data"] = data
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["flags"] = flags
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["priority"] = priority
            __props__.__dict__["tag"] = tag
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["weight"] = weight
            __props__.__dict__["domain_record"] = None
        super(DomainsRecordMx, __self__).__init__(
            'digitalocean-native:domains/v2:DomainsRecordMx',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DomainsRecordMx':
        """
        Get an existing DomainsRecordMx resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DomainsRecordMxArgs.__new__(DomainsRecordMxArgs)

        __props__.__dict__["data"] = None
        __props__.__dict__["domain_record"] = None
        __props__.__dict__["flags"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["priority"] = None
        __props__.__dict__["tag"] = None
        __props__.__dict__["ttl"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["weight"] = None
        return DomainsRecordMx(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[str]]:
        """
        Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="domainRecord")
    def domain_record(self) -> pulumi.Output[Optional['outputs.DomainRecord']]:
        return pulumi.get(self, "domain_record")

    @property
    @pulumi.getter
    def flags(self) -> pulumi.Output[Optional[int]]:
        """
        An unsigned integer between 0-255 used for CAA records.
        """
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The host name, alias, or service being defined by the record.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port for SRV records.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        The priority for SRV and MX records.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[Optional[str]]:
        """
        The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the DNS record. For example: A, CNAME, TXT, ...
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[Optional[int]]:
        """
        The weight for SRV records.
        """
        return pulumi.get(self, "weight")

