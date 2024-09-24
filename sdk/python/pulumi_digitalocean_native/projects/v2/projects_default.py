# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['ProjectsDefaultArgs', 'ProjectsDefault']

@pulumi.input_type
class ProjectsDefaultArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input['ProjectBaseEnvironment']] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[int]] = None,
                 owner_uuid: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectsDefault resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the project was created.
        :param pulumi.Input[str] description: The description of the project. The maximum length is 255 characters.
        :param pulumi.Input['ProjectBaseEnvironment'] environment: The environment of the project's resources.
        :param pulumi.Input[bool] is_default: If true, all resources will be added to this project if no project is specified.
        :param pulumi.Input[str] name: The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
        :param pulumi.Input[int] owner_id: The integer id of the project owner.
        :param pulumi.Input[str] owner_uuid: The unique universal identifier of the project owner.
        :param pulumi.Input[str] purpose: The purpose of the project. The maximum length is 255 characters. It can
               have one of the following values:
               
               - Just trying out DigitalOcean
               - Class project / Educational purposes
               - Website or blog
               - Web Application
               - Service or API
               - Mobile Application
               - Machine learning / AI / Data processing
               - IoT
               - Operational / Developer tooling
               
               If another value for purpose is specified, for example, "your custom purpose",
               your purpose will be stored as `Other: your custom purpose`.
        :param pulumi.Input[str] updated_at: A time value given in ISO8601 combined date and time format that represents when the project was updated.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if owner_uuid is not None:
            pulumi.set(__self__, "owner_uuid", owner_uuid)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the project. The maximum length is 255 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input['ProjectBaseEnvironment']]:
        """
        The environment of the project's resources.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input['ProjectBaseEnvironment']]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, all resources will be added to this project if no project is specified.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[int]]:
        """
        The integer id of the project owner.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="ownerUuid")
    def owner_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        The unique universal identifier of the project owner.
        """
        return pulumi.get(self, "owner_uuid")

    @owner_uuid.setter
    def owner_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_uuid", value)

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        The purpose of the project. The maximum length is 255 characters. It can
        have one of the following values:

        - Just trying out DigitalOcean
        - Class project / Educational purposes
        - Website or blog
        - Web Application
        - Service or API
        - Mobile Application
        - Machine learning / AI / Data processing
        - IoT
        - Operational / Developer tooling

        If another value for purpose is specified, for example, "your custom purpose",
        your purpose will be stored as `Other: your custom purpose`.
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ProjectsDefault(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input['ProjectBaseEnvironment']] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[int]] = None,
                 owner_uuid: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ProjectsDefault resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the project was created.
        :param pulumi.Input[str] description: The description of the project. The maximum length is 255 characters.
        :param pulumi.Input['ProjectBaseEnvironment'] environment: The environment of the project's resources.
        :param pulumi.Input[bool] is_default: If true, all resources will be added to this project if no project is specified.
        :param pulumi.Input[str] name: The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
        :param pulumi.Input[int] owner_id: The integer id of the project owner.
        :param pulumi.Input[str] owner_uuid: The unique universal identifier of the project owner.
        :param pulumi.Input[str] purpose: The purpose of the project. The maximum length is 255 characters. It can
               have one of the following values:
               
               - Just trying out DigitalOcean
               - Class project / Educational purposes
               - Website or blog
               - Web Application
               - Service or API
               - Mobile Application
               - Machine learning / AI / Data processing
               - IoT
               - Operational / Developer tooling
               
               If another value for purpose is specified, for example, "your custom purpose",
               your purpose will be stored as `Other: your custom purpose`.
        :param pulumi.Input[str] updated_at: A time value given in ISO8601 combined date and time format that represents when the project was updated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProjectsDefaultArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ProjectsDefault resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProjectsDefaultArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectsDefaultArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input['ProjectBaseEnvironment']] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[int]] = None,
                 owner_uuid: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectsDefaultArgs.__new__(ProjectsDefaultArgs)

            __props__.__dict__["created_at"] = created_at
            __props__.__dict__["description"] = description
            __props__.__dict__["environment"] = environment
            __props__.__dict__["is_default"] = is_default
            __props__.__dict__["name"] = name
            __props__.__dict__["owner_id"] = owner_id
            __props__.__dict__["owner_uuid"] = owner_uuid
            __props__.__dict__["purpose"] = purpose
            __props__.__dict__["updated_at"] = updated_at
        super(ProjectsDefault, __self__).__init__(
            'digitalocean-native:projects/v2:ProjectsDefault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProjectsDefault':
        """
        Get an existing ProjectsDefault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProjectsDefaultArgs.__new__(ProjectsDefaultArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["environment"] = None
        __props__.__dict__["is_default"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["owner_uuid"] = None
        __props__.__dict__["purpose"] = None
        __props__.__dict__["updated_at"] = None
        return ProjectsDefault(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the project. The maximum length is 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output['ProjectBaseEnvironment']:
        """
        The environment of the project's resources.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[bool]:
        """
        If true, all resources will be added to this project if no project is specified.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[Optional[int]]:
        """
        The integer id of the project owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="ownerUuid")
    def owner_uuid(self) -> pulumi.Output[Optional[str]]:
        """
        The unique universal identifier of the project owner.
        """
        return pulumi.get(self, "owner_uuid")

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Output[str]:
        """
        The purpose of the project. The maximum length is 255 characters. It can
        have one of the following values:

        - Just trying out DigitalOcean
        - Class project / Educational purposes
        - Website or blog
        - Web Application
        - Service or API
        - Mobile Application
        - Machine learning / AI / Data processing
        - IoT
        - Operational / Developer tooling

        If another value for purpose is specified, for example, "your custom purpose",
        your purpose will be stored as `Other: your custom purpose`.
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was updated.
        """
        return pulumi.get(self, "updated_at")

