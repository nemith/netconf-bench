{
  description = "NETCONF benchmark dev environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Build tools
            cmake
            pkg-config

            # Libraries
            openssl
            libssh

            # Python
            python314
            uv

            # Go
            go

            # Benchmarking
            hyperfine
          ];

          shellHook = ''
            export OPENSSL_ROOT_DIR="${pkgs.openssl.dev}"
            export PKG_CONFIG_PATH="${pkgs.openssl.dev}/lib/pkgconfig:${pkgs.libssh}/lib/pkgconfig:$PKG_CONFIG_PATH"
          '';
        };
      }
    );
}
