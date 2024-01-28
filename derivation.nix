{ pkgs }:
pkgs.stdenv.mkDerivation rec {
  pname = "nordvpn";
  version = "3.17.0";
  versionHash = "sha256-0SyDD55P6UmUKgD9h0/dPkJtVpRHj+rlSCyfUwZeu/E=";

  src = pkgs.fetchurl {
    url =
      "https://repo.nordvpn.com/deb/nordvpn/debian/pool/main/nordvpn_${version}_amd64.deb";
    sha256 = versionHash;
  };

  buildInputs = [ pkgs.libxml2 pkgs.libidn2 ];
  nativeBuildInputs = [ pkgs.dpkg pkgs.autoPatchelfHook pkgs.stdenv.cc.cc.lib ];

  dontConfigure = true;
  dontBuild = true;

  unpackPhase = ''
    runHook preUnpack
    dpkg --fsys-tarfile $src | tar --extract
    runHook postUnpack
  '';

  installPhase = ''
    runHook preInstall
    runHook postInstall
  '';

  meta = with pkgs.lib; {
    description = "CLI client for NordVPN";
    homepage = "https://nordvpn.com";
    license = pkgs.lib.licenses.gpl3;
  };
}
