{ inputs, self, system }:
{ config, lib, pkgs, ... }:
let cfg = config.programs.nordvpn;
in {
  options.programs.nordvpn = { enable = lib.mkEnableOption "hello service"; };

  config = lib.mkIf cfg.enable {
    environment = { systemPackages = [ self.packages.${system}.default ]; };
    systemd.services.hello = {
      wantedBy = [ "multi-user.target" ];
      serviceConfig.ExecStart =
        "${self.packages.${system}.default}/bin/hello -g'Hello, ${
          lib.escapeShellArg cfg.greeter
        }!'";
    };
  };
}
