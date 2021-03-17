{ buildGoModule, libpcap }:

buildGoModule rec {
  pname = "unchained";
  version = "0.0.1";

  src = builtins.filterSource (path: type: type != "directory" || baseNameOf path != ".git") ./.;

  vendorSha256 = "sha256:0c5qxhndh0cxin3vqdk02lpys525nxsmd9lf2mra7dylklbqfb2i"; 

  subPackages = [ "." ]; 

  runVend = true;

  buildInputs = [ libpcap ];
}


