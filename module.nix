{ inputs, self, system }:
{ config, lib, pkgs, ... }:
let cfg = config.programs.nordvpn;
in {
  options.programs.nordvpn = { enable = lib.mkEnableOption "nordvpn"; };

  config = lib.mkIf cfg.enable {
    environment = {
      systemPackages = [ "${self.packages.${system}.default}/bin/nordvpn" ];
    };
    systemd.services.nordvpnd = {
      wantedBy = [ "multi-user.target" ];
      serviceConfig.ExecStart =
        "${self.packages.${system}.default}/bin/nordvpnd";
    };
  };
}
