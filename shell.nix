{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.pkg-config
  ];

  buildInputs = [
    pkgs.go
    pkgs.mpv
  ];
}
