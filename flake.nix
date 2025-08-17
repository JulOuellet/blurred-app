{
  description = "Sportlight local development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        go
        air
        docker
        docker-compose
        postgresql
      ];

      shellHook = ''
        echo "----------------------------"
        echo "Development environment ready"
        echo "Go version: $(go version)"
        echo "Docker version: $(docker --version)"
        echo "Docker Compose version: $(docker-compose --version)"
        echo "Air version: $(air -v)"
        echo "----------------------------"
      '';
    };
  };
}
