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

__all__ = [
    'Account',
    'AccountTeamProperties',
    'GetAccountProperties',
    'GetSshKeysProperties',
    'ListSshKeys',
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
    'SshKeys',
]

@pulumi.output_type
class Account(dict):
    def __init__(__self__, *,
                 droplet_limit: int,
                 email: str,
                 email_verified: Optional[bool] = None,
                 floating_ip_limit: int,
                 status: Optional['AccountStatus'] = None,
                 status_message: str,
                 uuid: str,
                 name: Optional[str] = None,
                 team: Optional['outputs.AccountTeamProperties'] = None):
        """
        :param int droplet_limit: The total number of Droplets current user or team may have active at one time.
        :param str email: The email address used by the current user to register for DigitalOcean.
        :param bool email_verified: If true, the user has verified their account via email. False otherwise.
        :param int floating_ip_limit: The total number of Floating IPs the current user or team may have.
        :param 'AccountStatus' status: This value is one of "active", "warning" or "locked".
        :param str status_message: A human-readable message giving more details about the status of the account.
        :param str uuid: The unique universal identifier for the current user.
        :param str name: The display name for the current user.
        :param 'AccountTeamProperties' team: When authorized in a team context, includes information about the current team.
        """
        pulumi.set(__self__, "droplet_limit", droplet_limit)
        pulumi.set(__self__, "email", email)
        if email_verified is None:
            email_verified = False
        pulumi.set(__self__, "email_verified", email_verified)
        pulumi.set(__self__, "floating_ip_limit", floating_ip_limit)
        if status is None:
            status = 'active'
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "status_message", status_message)
        pulumi.set(__self__, "uuid", uuid)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if team is not None:
            pulumi.set(__self__, "team", team)

    @property
    @pulumi.getter(name="dropletLimit")
    def droplet_limit(self) -> int:
        """
        The total number of Droplets current user or team may have active at one time.
        """
        return pulumi.get(self, "droplet_limit")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email address used by the current user to register for DigitalOcean.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> bool:
        """
        If true, the user has verified their account via email. False otherwise.
        """
        return pulumi.get(self, "email_verified")

    @property
    @pulumi.getter(name="floatingIpLimit")
    def floating_ip_limit(self) -> int:
        """
        The total number of Floating IPs the current user or team may have.
        """
        return pulumi.get(self, "floating_ip_limit")

    @property
    @pulumi.getter
    def status(self) -> 'AccountStatus':
        """
        This value is one of "active", "warning" or "locked".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        A human-readable message giving more details about the status of the account.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        The unique universal identifier for the current user.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The display name for the current user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def team(self) -> Optional['outputs.AccountTeamProperties']:
        """
        When authorized in a team context, includes information about the current team.
        """
        return pulumi.get(self, "team")


@pulumi.output_type
class AccountTeamProperties(dict):
    """
    When authorized in a team context, includes information about the current team.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 uuid: Optional[str] = None):
        """
        When authorized in a team context, includes information about the current team.
        :param str name: The name for the current team.
        :param str uuid: The unique universal identifier for the current team.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name for the current team.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[str]:
        """
        The unique universal identifier for the current team.
        """
        return pulumi.get(self, "uuid")


@pulumi.output_type
class GetAccountProperties(dict):
    def __init__(__self__, *,
                 account: Optional['outputs.Account'] = None):
        if account is not None:
            pulumi.set(__self__, "account", account)

    @property
    @pulumi.getter
    def account(self) -> Optional['outputs.Account']:
        return pulumi.get(self, "account")


@pulumi.output_type
class GetSshKeysProperties(dict):
    def __init__(__self__, *,
                 ssh_key: Optional['outputs.SshKeys'] = None):
        if ssh_key is not None:
            pulumi.set(__self__, "ssh_key", ssh_key)

    @property
    @pulumi.getter(name="sshKey")
    def ssh_key(self) -> Optional['outputs.SshKeys']:
        return pulumi.get(self, "ssh_key")


@pulumi.output_type
class ListSshKeys(dict):
    def __init__(__self__, *,
                 meta: 'outputs.MetaMeta',
                 links: Optional['outputs.PageLinks'] = None,
                 ssh_keys: Optional[Sequence['outputs.SshKeys']] = None):
        pulumi.set(__self__, "meta", meta)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if ssh_keys is not None:
            pulumi.set(__self__, "ssh_keys", ssh_keys)

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter(name="sshKeys")
    def ssh_keys(self) -> Optional[Sequence['outputs.SshKeys']]:
        return pulumi.get(self, "ssh_keys")


@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[int] = None):
        """
        :param int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[int]:
        """
        Number of objects returned by the request.
        """
        return pulumi.get(self, "total")


@pulumi.output_type
class PageLinks(dict):
    def __init__(__self__, *,
                 pages: Optional['outputs.PageLinksPagesProperties'] = None):
        if pages is not None:
            pulumi.set(__self__, "pages", pages)

    @property
    @pulumi.getter
    def pages(self) -> Optional['outputs.PageLinksPagesProperties']:
        return pulumi.get(self, "pages")


@pulumi.output_type
class PageLinksPagesProperties(dict):
    def __init__(__self__, *,
                 first: Optional[str] = None,
                 last: Optional[str] = None,
                 next: Optional[str] = None,
                 prev: Optional[str] = None):
        if first is not None:
            pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if next is not None:
            pulumi.set(__self__, "next", next)
        if prev is not None:
            pulumi.set(__self__, "prev", prev)

    @property
    @pulumi.getter
    def first(self) -> Optional[str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[str]:
        return pulumi.get(self, "prev")


@pulumi.output_type
class SshKeys(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "publicKey":
            suggest = "public_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SshKeys. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SshKeys.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SshKeys.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 public_key: str,
                 fingerprint: Optional[str] = None,
                 id: Optional[int] = None):
        """
        :param str name: A human-readable display name for this key, used to easily identify the SSH keys when they are displayed.
        :param str public_key: The entire public key string that was uploaded. Embedded into the root user's `authorized_keys` file if you include this key during Droplet creation.
        :param str fingerprint: A unique identifier that differentiates this key from other keys using  a format that SSH recognizes. The fingerprint is created when the key is added to your account.
        :param int id: A unique identification number for this key. Can be used to embed a  specific SSH key into a Droplet.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "public_key", public_key)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-readable display name for this key, used to easily identify the SSH keys when they are displayed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> str:
        """
        The entire public key string that was uploaded. Embedded into the root user's `authorized_keys` file if you include this key during Droplet creation.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[str]:
        """
        A unique identifier that differentiates this key from other keys using  a format that SSH recognizes. The fingerprint is created when the key is added to your account.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        A unique identification number for this key. Can be used to embed a  specific SSH key into a Droplet.
        """
        return pulumi.get(self, "id")


