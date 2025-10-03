{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        package = pkgs.buildGoModule {
          pname = "backupmgr";
          version = "0.1.0";
          src = ./.;
          vendorHash = null;
          buildInputs = [];
        };
      in
      {
        packages = {
          default = package;
          backupmgr = package;
        };
      });
}
