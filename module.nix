{ inputs, self, system }:
{ config, lib, pkgs, ... }:
let
  cfg = config.programs.nordvpn;
  pkg = self.packages.${system}.default;
in {
  options.programs.nordvpn = { enable = lib.mkEnableOption "nordvpn"; };

  config = lib.mkIf cfg.enable {
    environment = {
      systemPackages = [
        pkg
        iproute2
        sysctl
        iptables
        procps
        cacert
        libxml2
        libidn2
        zlib
        wireguard-tools
      ];
    };
    users.groups.nordvpn = { };
    systemd.services.nordvpnd = {
      description = "NordVPN daemon.";
      serviceConfig = {
        Environment = "PATH=/run/current-system/sw/bin/";
        ExecStart = "${pkg}/bin/nordvpnd";
        ExecStartPre = ''
          ${pkgs.bash}/bin/bash -c '\
            mkdir -m 700 -p /var/lib/nordvpn; \
            cp -r ${pkg}/var/lib/nordvpn/* /var/lib/nordvpn; \
            '
        '';
        NonBlocking = true;
        KillMode = "process";
        Restart = "on-failure";
        RestartSec = 5;
        RuntimeDirectory = "nordvpn";
        RuntimeDirectoryMode = "0750";
        Group = "nordvpn";
      };
      wantedBy = [ "multi-user.target" ];
      after = [ "network-online.target" ];
      wants = [ "network-online.target" ];
    };
  };
}
