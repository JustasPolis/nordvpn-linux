{ inputs, self, system }:
{ config, lib, pkgs, ... }:
let
  cfg = config.programs.nordvpn;
  pkg = self.packages.${system}.default;
in {
  options.programs.nordvpn = { enable = lib.mkEnableOption "nordvpn"; };

  config = lib.mkIf cfg.enable {
    services.resolved = { enable = true; };
    networking.resolvconf.enable = true;
    environment = { systemPackages = [ pkg ]; };
    users.groups.nordvpn = { };
    systemd.services.nordvpnd = {
      path = with pkgs; [
        iproute2
        sysctl
        iptables
        procps
        cacert
        libxml2
        libidn2
        zlib
        wireguard-tools
        e2fsprogs
        systemd
      ];
      description = "NordVPN daemon.";
      wants = [ "network.target" ];
      after = [
        "network-online.target"
        "NetworkManager.service"
        "systemd-resolved.service"
      ];
      serviceConfig = {
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
    };
  };
}
