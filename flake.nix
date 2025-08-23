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
        export PATH="$HOME/go/bin:$PATH"

        if ! command -v migrate &> /dev/null; then
          echo "Installing golang-migrate..."
          go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
        fi


        echo "----------------------------"
        echo "Development environment ready"
        echo "Go version: $(go version)"
        echo "Docker version: $(docker --version)"
        echo "Docker Compose version: $(docker-compose --version)"
        echo "----------------------------"
      '';
    };
  };
}
