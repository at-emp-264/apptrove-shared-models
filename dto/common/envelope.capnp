@0x8424671d2d0d7366;

using Go = import "/go.capnp";
$Go.package("dto");
$Go.import("github.com/ApptroveLabs/apptrove-shared-models/dto/common");

struct Header {
  key   @0 :Text;
  value @1 :Text;
}

struct Envelope {
  id      @0 :Text;
  reqTs   @1 :Int64;
  headers @2 :List(Header);
  type    @3 :Text;
}