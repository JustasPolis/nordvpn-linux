{ pkgs }:

pkgs.buildGo120Module rec {
  pname = "nordvpn";
  version = "3.17.0";
  versionHash = "sha256-ONdZIhJZmtqxwGeKHQ2LuaIWj+QcCmQy9iARhbLQzWI=";

  proxyVendor = true;

  nativeBuildInputs = [ pkgs.pkg-config ];

  buildInputs = [ pkgs.libxml2.dev pkgs.iptables pkgs.iproute2 ];

  ldflags = [
    "-X main.Version=${version}"
    "-X main.Salt=${"abc123"}"
    "-X main.Environment=${"prod"}"
    "-X main.Hash=${versionHash}"
  ];

  CGO_ENABLED = 1;

  CGO_CFLAGS = [ "-g" "-O2" "-D_FORTIFY_SOURCE=2" ];
  CGO_LDFLAGS = [ "-Wl,-z,relro,-z,now" ];

  src = pkgs.fetchFromGitHub {
    owner = "JustasPolis";
    repo = "nordvpn-linux";
    rev = version;
    sha256 = versionHash;
  };

  vendorHash = "sha256-IJpvL4Q+uvEFLgyls1tGa8SD40y+SRnnfk++/yfiWE0=";

  meta = with pkgs.lib; {
    description = "CLI client for NordVPN";
    homepage = "https://nordvpn.com";
    license = pkgs.lib.licenses.gpl3;
  };
}
