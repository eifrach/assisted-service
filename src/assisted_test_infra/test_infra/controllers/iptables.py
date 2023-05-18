from enum import Enum
from ipaddress import IPv4Address, IPv6Address
from typing import List, Optional

from assisted_test_infra.test_infra import utils
from service_client import log


class IpTableCommandOption(Enum):
    CHECK = "check"
    INSERT = "insert"
    DELETE = "delete"


class IptableRule:
    CHAIN_INPUT = "INPUT"
    CHAIN_FORWARD = "FORWARD"

    def __init__(
        self,
        chain: str,
        target: str,
        protocol: str,
        dest_port: Optional[str] = "",
        sources: Optional[List] = None,
        extra_args: Optional[str] = "",
        address_familiy=IPv4Address,
    ):
        self._chain = chain
        self._target = target
        self._protocol = protocol
        self._dest_port = dest_port
        self._sources = sources if sources else []
        self._extra_args = extra_args
        self.address_familiy = address_familiy

    @property
    def _iptables_bin(self):
        return "ip6tables" if self.address_familiy is IPv6Address else "iptables"

    def _build_command_string(self, option: IpTableCommandOption) -> str:
        sources_string = ",".join(self._sources)
        rule_template = [
            f"{self._iptables_bin}",
            f"--{option.value}",
            self._chain,
            "-p",
            self._protocol,
            "-j",
            self._target,
        ]

        if self._sources:
            rule_template += ["-s", sources_string]

        if self._dest_port:
            rule_template += ["--dport", self._dest_port]

        if self._extra_args:
            rule_template += [self._extra_args]
        rule_str = " ".join(rule_template)
        log.info(f"build iptables command: {rule_str}")
        return rule_str

    def _does_rule_exist(self) -> bool:
        check_rule = self._build_command_string(IpTableCommandOption.CHECK)
        _, _, exit_code = utils.run_command(check_rule, shell=True, raise_errors=False)

        return exit_code == 0

    def add_sources(self, sources):
        self._sources += sources

    def insert(self) -> None:
        if not self._does_rule_exist():
            insert_rule = self._build_command_string(IpTableCommandOption.INSERT)
            utils.run_command(insert_rule, shell=True)

    def delete(self) -> None:
        if self._does_rule_exist():
            delete_rule = self._build_command_string(IpTableCommandOption.DELETE)
            utils.run_command(delete_rule, shell=True)
